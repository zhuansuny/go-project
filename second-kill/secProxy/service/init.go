package service

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

var (
	secKillConf *SecSkillConf //全局变量
)

func InitService(serviceConf *SecSkillConf) (err error) {
	secKillConf = serviceConf
	//从redis加载黑名单
	err = loadBlackList()
	if err != nil {
		logs.Error("load black list err:%v", err)
		return
	}
	logs.Debug("init service succ, config:%v", secKillConf)

	err = initProxy2LayerRedis()
	if err != nil {
		logs.Error("load proxy2layer redis pool failed, err:%v", err)
		return
	}

	secKillConf.secLimitMgr = &SecLimitMgr{
		UserLimitMap: make(map[int]*Limit, 1000),
		IpLimitMap:   make(map[string]*Limit, 1000),
	}

	secKillConf.SecReqChan = make(chan *SecRequest, 1000) //secKillConf.SecReqChanSize
	secKillConf.UserConnMap = make(map[string]chan *SecResult, 1000)

	initRedisProcessFunc()

	return
}

//开启多个（默认16个）协程从redis中读以及写数据
func initRedisProcessFunc() {
	for i := 0; i < secKillConf.WriteProxy2LayerGoroutineNum; i++ {
		go WriteHandle()
	}

	for i := 0; i < secKillConf.ReadProxy2LayerGoroutineNum; i++ {
		go ReadHandle()
	}
}

//初始化数据库连接池（业务逻辑层到接入层）
func initProxy2LayerRedis() (err error) {
	secKillConf.proxy2LayerRedisPool = &redis.Pool{
		MaxIdle:     secKillConf.RedisProxy2LayerConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisProxy2LayerConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisProxy2LayerConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.RedisProxy2LayerConf.RedisAddr)
		},
	}

	conn := secKillConf.proxy2LayerRedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}

	return
}

//初始化数据库连接池（接入层到业务逻辑层）
// func initLayer2ProxyRedis() (err error) {
// 	secKillConf.layer2ProxyRedisPool = &redis.Pool{
// 		MaxIdle:     secKillConf.RedisLayer2ProxyConf.RedisMaxIdle,
// 		MaxActive:   secKillConf.RedisLayer2ProxyConf.RedisMaxActive,
// 		IdleTimeout: time.Duration(secKillConf.RedisLayer2ProxyConf.RedisIdleTimeout) * time.Second,
// 		Dial: func() (redis.Conn, error) {
// 			return redis.Dial("tcp", secKillConf.RedisLayer2ProxyConf.RedisAddr)
// 		},
// 	}

// 	conn := secKillConf.layer2ProxyRedisPool.Get()
// 	defer conn.Close()

// 	_, err = conn.Do("ping")
// 	if err != nil {
// 		logs.Error("ping redis failed, err:%v", err)
// 		return
// 	}

// 	return
// }

//初始化黑名单数控库连接池
func initBlackRedis() (err error) {
	secKillConf.blackRedisPool = &redis.Pool{
		MaxIdle:     secKillConf.RedisBlackConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisBlackConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.RedisBlackConf.RedisAddr)
		},
	}

	conn := secKillConf.blackRedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}

	return
}

//导入黑名单
func loadBlackList() (err error) {
	//线make分配地址
	secKillConf.ipBlackMap = make(map[string]bool, 1000)
	secKillConf.idBlackMap = make(map[int]bool, 1000)
	//初始化连接池
	err = initBlackRedis()
	if err != nil {
		logs.Error("init black redis failed, err:%v", err)
		return
	}

	conn := secKillConf.blackRedisPool.Get() //获取连接
	defer conn.Close()

	reply, err := conn.Do("hgetall", "idblacklist")
	idlist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err:%v", err)
		return
	}

	for _, v := range idlist {
		id, err := strconv.Atoi(v)
		if err != nil {
			logs.Warn("invalid user id [%v]", id)
			continue
		}
		secKillConf.idBlackMap[id] = true
	}

	reply, err = conn.Do("hgetall", "ipblacklist")
	iplist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err:%v", err)
		return
	}

	for _, v := range iplist {
		secKillConf.ipBlackMap[v] = true
	}
	//开启协程异步处理新加入的黑名单
	go SyncIpBlackList()
	go SyncIdBlackList()
	return
}

func SyncIpBlackList() {
	var ipList []string
	lastTime := time.Now().Unix()
	for {
		conn := secKillConf.blackRedisPool.Get()
		defer conn.Close()
		reply, err := conn.Do("BLPOP", "blackiplist", 0) //阻塞一秒
		ip, err := redis.String(reply, err)
		if err != nil {
			continue
		}

		curTime := time.Now().Unix()
		ipList = append(ipList, ip)

		if len(ipList) > 100 || curTime-lastTime > 5 { //黑名单数目加入100个，或者每过5秒写入黑名单
			secKillConf.RWBlackLock.Lock()
			for _, v := range ipList {
				secKillConf.ipBlackMap[v] = true
			}
			secKillConf.RWBlackLock.Unlock()

			lastTime = curTime
			logs.Info("sync ip list from redis succ, ip[%v]", ipList)
		}
	}
}

func SyncIdBlackList() {
	for {
		conn := secKillConf.blackRedisPool.Get()
		defer conn.Close()
		reply, err := conn.Do("BLPOP", "blackidlist", 0) //阻塞一秒
		id, err := redis.Int(reply, err)
		if err != nil {
			continue
		}

		secKillConf.RWBlackLock.Lock()
		secKillConf.idBlackMap[id] = true
		secKillConf.RWBlackLock.Unlock()

		logs.Info("sync id list from redis succ, ip[%v]", id)
	}
}
