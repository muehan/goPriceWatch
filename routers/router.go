package routers

import (
	"goPriceWatch/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/product", &controllers.ProductController{})
}
