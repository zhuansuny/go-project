package controllers

import (
	//"fmt"
	"secKill/secProxy/service"
	//"strconv"
	//"strings"
	//"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type SkillController struct {
	beego.Controller
}

//秒杀页面逻辑处理
func (p *SkillController) SecKill() {
	productId, err := p.GetInt("product_id")
	result := make(map[string]interface{})

	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON()
	}()

	if err != nil {
		result["code"] = 1001
		result["message"] = "invalid product_id"
		return
	}

	// source := p.GetString("src")        //客户端来源
	// authcode := p.GetString("authcode") //生成的校验码
	secTime := p.GetString("time") //客户端时间
	// nance := p.GetString("nance")       //随机码

	secRequest := service.NewSecRequest()
	secRequest.ResultChan = make(chan *service.SecResult, 1)
	//secRequest.AuthCode = authcode
	//secRequest.Nance = nance
	secRequest.ProductId = productId
	secRequest.SecTime = secTime
	//secRequest.Source = source
	// secRequest.UserAuthSign = p.Ctx.GetCookie("userAuthSign")
	// secRequest.UserId, _ = strconv.Atoi(p.Ctx.GetCookie("userId"))
	// secRequest.AccessTime = time.Now()
	// if len(p.Ctx.Request.RemoteAddr) > 0 {
	// 	secRequest.ClientAddr = strings.Split(p.Ctx.Request.RemoteAddr, ":")[0] //获取用户IP
	// }

	secRequest.UserId, err = p.GetInt("user_id")
	if err != nil {
		result["code"] = 1001
		result["message"] = "invalid user_id"
		return
	}

	secRequest.ClientRefence = p.Ctx.Request.Referer()
	secRequest.CloseNotify = p.Ctx.ResponseWriter.CloseNotify()

	logs.Debug("client request:[%v]", secRequest)
	// if err != nil {
	// 	result["code"] = service.ErrInvalidRequest
	// 	result["message"] = fmt.Sprintf("invalid cookie:userId")
	// 	return
	// }

	data, code, err := service.SecKill(secRequest)
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()
		return
	}

	result["data"] = data
	result["code"] = code

	return
}

//秒杀商品逻辑处理
func (p *SkillController) SecInfo() {

	productId, err := p.GetInt("product_id")
	result := make(map[string]interface{})

	result["code"] = 0
	result["message"] = "success"

	defer func() {
		p.Data["json"] = result
		p.ServeJSON() //以json数据展示
	}()

	if err != nil {
		result["code"] = 1001
		result["message"] = "invalid product_id"

		logs.Error("invalid request, get product_id failed, err:%v", err)
		return
	}
	logs.Debug("productId:", productId)
	//data, code, err := service.SecInfo(productId)
	data, code, err := service.SecInfoList()
	if err != nil {
		result["code"] = code
		result["message"] = err.Error()

		logs.Error("invalid request, get product_id failed, err:%v", err)
		return
	}

	result["data"] = data
}
