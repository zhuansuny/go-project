package routers

import (
	"secKill/secAdmin/controllers/activity"
	"secKill/secAdmin/controllers/product"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/product/list", &product.ProductController{}, "*:ListProduct")
	beego.Router("/", &product.ProductController{}, "*:ListProduct")
	beego.Router("/product/create", &product.ProductController{}, "*:CreateProduct")
	beego.Router("/product/apply", &product.ProductController{}, "*:ApplyProduct")

	beego.Router("/activity/create", &activity.ActivityController{}, "*:CreateActivity")
	beego.Router("/activity/list", &activity.ActivityController{}, "*:ListActivity")
	beego.Router("/activity/apply", &activity.ActivityController{}, "*:ApplyActivity")
}
