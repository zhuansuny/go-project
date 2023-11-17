package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
)

func WriteHandle() {

	for {
		logs.Debug("begin get from secReqchan")
		req := <-secKillConf.SecReqChan //从channel中获取到一个请求，用户抢购的秒杀商品信息
		logs.Debug("success get from secReqchan")
		conn := secKillConf.proxy2LayerRedisPool.Get() //从redis池中获取连接

		data, err := json.Marshal(req)
		if err != nil {
			logs.Error("json.Marshal failed, error:%v req:%v", err, req)
			conn.Close()
			continue
		}
		_, err = conn.Do("LPUSH", "sec_queue", data)
		if err != nil {
			logs.Error("lpush failed, err:%v, req:%v", err, req)
			conn.Close()
			continue
		}
		logs.Debug("success,发送到逻辑层", data)
		conn.Close()
	}

}

func ReadHandle() {
	for {
		conn := secKillConf.proxy2LayerRedisPool.Get()

		reply, err := conn.Do("RPOP", "recv_queue") //读取最终秒杀成功的用户信息（从逻辑层处理完发送回来）
		data, err := redis.String(reply, err)
		//reply, err := conn.Do("blpop", "recv_queue", 0)

		if err == redis.ErrNil { //若没查到数据，休眠一秒进行下一循环
			time.Sleep(time.Second)
			conn.Close()
			continue
		}
		if err != nil {
			logs.Error("rpop failed, err:%v", err)
			conn.Close()
			continue
		}
		logs.Debug("rpop from redis succ, data:%s", data)
		var result SecResult
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			logs.Error("json.Unmarshal failed, err:%v", err)
			conn.Close()
			continue
		}

		userKey := fmt.Sprintf("%s_%s", result.UserId, result.ProductId) //将反序列化之后的UserId和ProductId拼接

		secKillConf.UserConnMapLock.Lock()
		resultChan, ok := secKillConf.UserConnMap[userKey] //从map中获取用户商品信息中的result通道
		secKillConf.UserConnMapLock.Unlock()
		if !ok {
			conn.Close()
			logs.Warn("user not found:%v", userKey)
			continue
		}

		resultChan <- &result //将result放入通道，（service中的通道堵塞消失，代表秒杀商品成功）
		conn.Close()
	}
}
