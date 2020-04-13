package routers

import (
	"goPriceWatch/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// beego.Router("/product", &controllers.ProductController{}, "get:GetProducts")
	// beego.Router("/product/:id", &controllers.ProductController{},  "get:Details")
	// beego.Router("/api/product", &controllers.ProductApiController{})
	// beego.Router("/api/product/:id", &controllers.ProductApiController{},  "get:GetDetails")
	// beego.Router("/api/product/:search", &controllers.ProductApiController{},  "get:Search")
	// beego.Router("/api/producttype", &controllers.ProductTypeApiController{})

	 restfulRouter := beego.NewNamespace("/api",
		 beego.NSNamespace("/product",
			beego.NSInclude(
			     &controllers.ProductController{},
		    ),
		),
		beego.NSNamespace("/producttype",
			beego.NSInclude(
				&controllers.ProductTypeController{},
			),
		 ),
		 beego.NSNamespace("/productbytype",
			beego.NSInclude(
				&controllers.ProductByTypeController{},
			),
		 ),
		 beego.NSNamespace("/price",
			beego.NSInclude(
				&controllers.PriceController{},
			),
	 	),
	 )

	 beego.AddNamespace(restfulRouter)
}
