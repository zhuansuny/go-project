package IndexController

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type IndexController struct {
	beego.Controller
}

func (p *IndexController) Index() {

	logs.Debug("enter index controller")
	//p.TplName = "index.html" 以HTML方式展现

	m := make(map[string]interface{})
	m["code"] = 200
	m["message"] = "success"
	p.Data["json"] = m

	p.ServeJSON(true) //以json数据展示

}
