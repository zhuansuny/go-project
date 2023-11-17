package service

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"

	"crypto/md5"
	"math/rand"
)

//初始化redis连接池
func initRedisPool(redisConf RedisConf) (pool *redis.Pool, err error) {
	pool = &redis.Pool{
		MaxIdle:     redisConf.RedisMaxIdle, //从传入的RedisConf实例读取配置
		MaxActive:   redisConf.RedisMaxActive,
		IdleTimeout: time.Duration(redisConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisConf.RedisAddr)
		},
	}

	conn := pool.Get()
	defer conn.Close()

	_, err = conn.Do("ping") //测试连接
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}
	return
}

//初始化连接
func initRedis(conf *SecLayerConf) (err error) {
	//创建接入层到业务逻辑层的连接池
	secLayerContext.proxy2LayerRedisPool, err = initRedisPool(conf.Proxy2LayerRedis)
	if err != nil {
		logs.Error("init proxy2layer redis pool failed, err:%v", err)
		return
	}
	//创建业务逻辑层到接入层的连接池
	secLayerContext.layer2ProxyRedisPool, err = initRedisPool(conf.Layer2ProxyRedis)
	if err != nil {
		logs.Error("init layer2proxy redis pool failed, err:%v", err)
		return
	}

	return
}

func RunProcess() (err error) {
	//开启设定数目的协程，来进行读写操作
	for i := 0; i < secLayerContext.secLayerConf.ReadGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleReader()
	}

	for i := 0; i < secLayerContext.secLayerConf.WriteGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleWrite()
	}
	//开启设定数目的协程，来进行处理用户业务操作，处理完，再放入写线程
	for i := 0; i < secLayerContext.secLayerConf.HandleUserGoroutineNum; i++ {
		secLayerContext.waitGroup.Add(1)
		go HandleUser()
	}

	logs.Debug("all process goroutine started")
	secLayerContext.waitGroup.Wait()
	logs.Debug("wait all goroutine exited")
	return
}

//进行读取从接入层发送的request数据，并加入管道准备让HandleUser处理
func HandleReader() {

	logs.Debug("read goroutine running")
	for {
		conn := secLayerContext.proxy2LayerRedisPool.Get() //从连接池获取连接
		for {
			//阻塞的从redis中取出数据
			ret, err := conn.Do("blpop", secLayerContext.secLayerConf.Proxy2LayerRedis.RedisQueueName, 0)
			if err != nil {
				logs.Error("pop from queue failed, err:%v", err)
				break
			}

			tmp, ok := ret.([]interface{})
			if !ok || len(tmp) != 2 {
				logs.Error("pop from queue failed, err:%v", err)
				continue
			}

			data, ok := tmp[1].([]byte)
			if !ok {
				logs.Error("pop from queue failed, err:%v", err)
				continue
			}

			logs.Debug("pop from queue, data:%s", string(data))

			var req SecRequest
			err = json.Unmarshal([]byte(data), &req)
			if err != nil {
				logs.Error("unmarshal to secrequest failed, err:%v", err)
				continue
			}
			//判断请求是否超时，超时直接返回（配置默认是30秒）
			// now := time.Now().Unix()
			// logs.Debug("req time:%v,now:%v", req.AccessTime.Unix(), now)
			// if now-req.AccessTime.Unix() >= int64(secLayerContext.secLayerConf.MaxRequestWaitTimeout) {
			// 	logs.Warn("req[%v] is expire", req)
			// 	continue
			// }

			//设置发送到管道的超时时间（配置默认100毫秒）
			timer := time.NewTicker(time.Millisecond * time.Duration(secLayerContext.secLayerConf.SendToHandleChanTimeout))
			select {
			case secLayerContext.Read2HandleChan <- &req:
			case <-timer.C:
				logs.Warn("send to handle chan timeout, req:%v", req)
				break
			}
		}

		conn.Close()
	}
}

//将从HandleUser处理后通过的用户发回给接入层redis,通知接入层
func HandleWrite() {
	logs.Debug("handle write running")
	//不断从channel中取出通过的数据
	for res := range secLayerContext.Handle2WriteChan {
		logs.Debug("begin send proxy:%v", res)
		err := sendToRedis(res) //发送给redis
		if err != nil {
			logs.Error("send to redis, err:%v, res:%v", err, res)
			continue
		}
	}
}

//发送数据给接入层redis
func sendToRedis(res *SecResponse) (err error) {

	data, err := json.Marshal(res) //将SecResponse序列化
	if err != nil {
		logs.Error("marshal failed, err:%v", err)
		return
	}

	conn := secLayerContext.layer2ProxyRedisPool.Get() //连接池取出连接
	_, err = conn.Do("rpush", secLayerContext.secLayerConf.Layer2ProxyRedis.RedisQueueName, string(data))
	if err != nil {
		logs.Warn("rpush to redis failed, err:%v", err)
		return
	}
	logs.Debug("success send proxy:%v", res)
	conn.Close()
	return
}

//秒杀业务核心处理
func HandleUser() {

	logs.Debug("handle user running")
	//从管道取出request
	for req := range secLayerContext.Read2HandleChan {
		logs.Debug("begin process request:%v", req)
		res, err := HandleSecKill(req)
		if err != nil {
			logs.Warn("process request %v failed, err:%v", err)
			res = &SecResponse{
				Code: ErrServiceBusy,
			}
		}
		//设置业务逻辑处理的超时时间（配置默认100毫秒）
		logs.Debug("request success:%v", res)
		timer := time.NewTicker(time.Millisecond * time.Duration(secLayerContext.secLayerConf.SendToWriteChanTimeout))

		select {
		case secLayerContext.Handle2WriteChan <- res:
		case <-timer.C:
			logs.Warn("send to response chan timeout, res:%v", res)
			break
		}

	}
	return
}

func HandleSecKill(req *SecRequest) (res *SecResponse, err error) {

	secLayerContext.RWSecProductLock.RLock() //读写锁
	defer secLayerContext.RWSecProductLock.RUnlock()

	res = &SecResponse{}
	res.UserId = req.UserId
	res.ProductId = req.ProductId
	//判断商品是否存在
	product, ok := secLayerContext.secLayerConf.SecProductInfoMap[req.ProductId]
	if !ok {
		logs.Error("not found product:%v", req.ProductId)
		res.Code = ErrNotFoundProduct
		return
	}
	//判断商品状态
	if product.Status == ProductStatusSoldout {
		res.Code = ErrSoldout
		return
	}
	//判断每秒钟抢购数量是否达到限制
	now := time.Now().Unix()
	alreadySoldCount := product.secLimit.Check(now)
	if alreadySoldCount >= product.SoldMaxLimit {
		res.Code = ErrRetry
		return
	}
	//判断是否已经购买
	secLayerContext.HistoryMapLock.Lock()
	userHistory, ok := secLayerContext.HistoryMap[req.UserId]
	if !ok {
		userHistory = &UserBuyHistory{
			history: make(map[int]int, 16),
		}

		secLayerContext.HistoryMap[req.UserId] = userHistory
	}

	histryCount := userHistory.GetProductBuyCount(req.ProductId)
	secLayerContext.HistoryMapLock.Unlock()

	if histryCount >= product.OnePersonBuyLimit {
		res.Code = ErrAlreadyBuy
		return
	}
	//判断商品是否秒杀完
	curSoldCount := secLayerContext.productCountMgr.Count(req.ProductId) //已经卖出的商品数量
	if curSoldCount >= product.Total {
		res.Code = ErrSoldout
		product.Status = ProductStatusSoldout
		return
	}

	//设定购买概率
	curRate := rand.Float64()
	fmt.Printf("curRate:%v product:%v count:%v total:%v\n", curRate, product.BuyRate, curSoldCount, product.Total)
	if curRate > product.BuyRate {
		res.Code = ErrRetry
		return
	}
	//历史秒杀记录，以及卖出数据加一
	userHistory.Add(req.ProductId, 1)
	secLayerContext.productCountMgr.Add(req.ProductId, 1)

	//用户id&商品id&当前时间&密钥
	res.Code = ErrSecKillSucc
	tokenData := fmt.Sprintf("userId=%d&productId=%d&timestamp=%d&security=%s",
		req.UserId, req.ProductId, now, secLayerContext.secLayerConf.TokenPasswd)

	res.Token = fmt.Sprintf("%x", md5.Sum([]byte(tokenData))) //返回数据data的MD5校验和。
	res.TokenTime = now

	return
}
