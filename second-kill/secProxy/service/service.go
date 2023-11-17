package service

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

func NewSecRequest() (secRequest *SecRequest) {
	secRequest = &SecRequest{
		ResultChan: make(chan *SecResult, 1),
	}

	return
}

func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {

	item, code, err := SecInfoById(productId)
	if err != nil {
		return
	}

	data = append(data, item)
	return
}

func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {

	secKillConf.RWSecProductLock.RLock()         //加锁
	defer secKillConf.RWSecProductLock.RUnlock() //延时解锁

	v, ok := secKillConf.SecProductInfoMap[productId] //获取对应id的商品信息
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found product_id:%d", productId)
		return
	}

	start := false
	end := false
	status := "success"

	now := time.Now().Unix() //时间戳
	if now-v.StartTime < 0 { //未到秒杀开始时间
		start = false
		end = false
		status = "sec kill is not start"
		code = ErrActiveNotStart
	}

	if now-v.StartTime > 0 { //秒杀开始
		start = true
	}

	if now-v.EndTime > 0 { //秒杀结束
		start = false
		end = true
		status = "sec kill is already end"
		code = ErrActiveAlreadyEnd
	}

	if v.Status == ProductStatusForceSaldOut || v.Status == ProductStatusSaldOut { //秒杀售罄或异常停止，秒杀结束
		start = false
		end = true
		status = "product is sald out"
		code = ErrActiveSaleOut
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start"] = start
	data["end"] = end
	data["status"] = status

	return
}

//获取秒杀商品配置列表
func SecInfoList() (data []map[string]interface{}, code int, err error) {

	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	for _, v := range secKillConf.SecProductInfoMap {

		item, _, err := SecInfoById(v.ProductId)
		if err != nil {
			logs.Error("get product_id[%d] failed, err:%v", v.ProductId, err)
			continue
		}

		logs.Debug("get product[%d]， result[%v], all[%v] v[%v]", v.ProductId, item, secKillConf.SecProductInfoMap, v)
		data = append(data, item)
	}

	return
}

func userCheck(req *SecRequest) (err error) {

	found := false
	//查看用户是否在白名单中
	for _, refer := range secKillConf.ReferWhiteList {
		if refer == req.ClientRefence {
			found = true
			break
		}
	}
	//不在就返回
	if !found {
		err = fmt.Errorf("invalid request")
		logs.Warn("user[%d] is reject by refer, req[%v]", req.UserId, req)
		return
	}
	//在就检查密钥
	authData := fmt.Sprintf("%d:%s", req.UserId, secKillConf.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData))) //将userId与密钥用md5加密，返回结果

	if authSign != req.UserAuthSign { //对比客户cookie中数据对比
		err = fmt.Errorf("invalid user cookie auth")
		return
	}
	return
}

func SecKill(req *SecRequest) (data map[string]interface{}, code int, err error) {

	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	// err = userCheck(req) //用户校验
	// if err != nil {
	// 	code = ErrUserCheckAuthFailed
	// 	logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
	// 	return
	// }

	// err = antiSpam(req) //反作弊校验（根据一定时间内访问次数）
	// if err != nil {
	// 	code = ErrUserServiceBusy
	// 	logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
	// 	return
	// }

	data, code, err = SecInfoById(req.ProductId) //根据商品id获取商品信息
	if err != nil {
		logs.Warn("userId[%d] secInfoBy Id failed, req[%v]", req.UserId, req)
		return
	}

	if code != 0 { //code不为0代表未抢到
		logs.Warn("userId[%d] secInfoByid failed, code[%d] req[%v]", req.UserId, code, req)
		return
	}

	userKey := fmt.Sprintf("%s_%s", req.UserId, req.ProductId) //抢到后将用户名称与商品id拼接加入Map，在线用户
	logs.Debug("userKey:", userKey)
	secKillConf.UserConnMap[userKey] = req.ResultChan

	secKillConf.SecReqChan <- req //将req加入到SecReqChan通道中,准备redis发送到逻辑层处理
	logs.Debug("insert secReqChan sueccss", userKey)
	ticker := time.NewTicker(time.Second * 10) //每隔十秒向ticher通道发送时间,代表十秒为超时时间

	defer func() {
		ticker.Stop() //关闭通道，释放资源
		secKillConf.UserConnMapLock.Lock()
		delete(secKillConf.UserConnMap, userKey)
		secKillConf.UserConnMapLock.Unlock()
	}()

	select {
	case <-ticker.C: //十秒后，报超时
		code = ErrProcessTimeout
		err = fmt.Errorf("request timeout")

		return
	case <-req.CloseNotify: //客户断开连接
		code = ErrClientClosed
		err = fmt.Errorf("client already closed")
		return
	case result := <-req.ResultChan: //秒杀商品成功
		code = result.Code
		data["product_id"] = result.ProductId
		data["token"] = result.Token
		data["user_id"] = result.UserId
		logs.Debug("%v抢购成功%v", req.UserId, req.ProductId)
		return
	}

	return
}
