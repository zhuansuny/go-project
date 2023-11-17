package service

import (
	"fmt"
	"sync"

	"github.com/astaxie/beego/logs"
)

type SecLimitMgr struct {
	UserLimitMap map[int]*Limit
	IpLimitMap   map[string]*Limit
	lock         sync.Mutex
}

func antiSpam(req *SecRequest) (err error) { //反作弊
	//检查是否在id,ip黑名单之中，是就返回错误
	_, ok := secKillConf.idBlackMap[req.UserId]
	if ok {
		err = fmt.Errorf("invalid request")
		logs.Error("useId[%v] is block by id black", req.UserId)
		return
	}

	_, ok = secKillConf.ipBlackMap[req.ClientAddr]
	if ok {
		err = fmt.Errorf("invalid request")
		logs.Error("useId[%v] ip[%v] is block by ip black", req.UserId, req.ClientAddr)
		return
	}

	secKillConf.secLimitMgr.lock.Lock()
	//根据userid 频率控制
	limit, ok := secKillConf.secLimitMgr.UserLimitMap[req.UserId] //判断该userid是否是首次访问，
	if !ok {                                                      //是就将加入频率计数map，开始计算访问频率
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		secKillConf.secLimitMgr.UserLimitMap[req.UserId] = limit
	}

	secIdCount := limit.secLimit.Count(req.AccessTime.Unix()) //把请求秒数传入判断，uid一秒内有几次访问
	minIdCount := limit.minLimit.Count(req.AccessTime.Unix()) //把请求秒数传入判断，uid一分钟内有几次访问

	//根据ip 频率控制
	limit, ok = secKillConf.secLimitMgr.IpLimitMap[req.ClientAddr] //判断该IP是否是首次访问
	if !ok {                                                       //是就加入频率计数map,开始计算访问频率
		limit = &Limit{
			secLimit: &SecLimit{},
			minLimit: &MinLimit{},
		}
		secKillConf.secLimitMgr.IpLimitMap[req.ClientAddr] = limit
	}

	secIpCount := limit.secLimit.Count(req.AccessTime.Unix()) //把请求秒数传入判断ip一秒内有几次访问
	minIpCount := limit.minLimit.Count(req.AccessTime.Unix()) //把请求秒数传入判断ip一分钟内有几次访问
	secKillConf.secLimitMgr.lock.Unlock()

	//---------访问次数超过设定值，就直接返回错误，不能再进行抢购----------
	if secIpCount > secKillConf.AccessLimitConf.IPSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if minIpCount > secKillConf.AccessLimitConf.IPMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if secIdCount > secKillConf.AccessLimitConf.UserSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	if minIdCount > secKillConf.AccessLimitConf.UserMinAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}
	return
}
