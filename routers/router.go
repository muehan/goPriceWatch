package routers

import (
	"goPriceWatch/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/product", &controllers.ProductController{}, "get:GetProducts")
	beego.Router("/product/:id", &controllers.ProductController{},  "get:Details")
	beego.Router("/api/product", &controllers.ProductApiController{})
	beego.Router("/api/product/:id", &controllers.ProductApiController{},  "get:GetDetails")
	beego.Router("/api/producttype", &controllers.ProductTypeApiController{})
}
