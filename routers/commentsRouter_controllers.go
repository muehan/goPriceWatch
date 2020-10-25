package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["goPriceWatch/controllers:PriceController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:PriceController"],
        beego.ControllerComments{
            Method: "GetByDate",
            Router: `/:date`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goPriceWatch/controllers:PriceController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:PriceController"],
        beego.ControllerComments{
            Method: "GetForProduct",
            Router: `/:productId/:days`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goPriceWatch/controllers:ProductByDateController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:ProductByDateController"],
        beego.ControllerComments{
            Method: "GetByDate",
            Router: `/:date`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goPriceWatch/controllers:ProductByTypeController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:ProductByTypeController"],
        beego.ControllerComments{
            Method: "GetProductsByType",
            Router: `/:productTypeId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goPriceWatch/controllers:ProductController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:ProductController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goPriceWatch/controllers:ProductController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:ProductController"],
        beego.ControllerComments{
            Method: "Search",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goPriceWatch/controllers:ProductController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:ProductController"],
        beego.ControllerComments{
            Method: "GetDetails",
            Router: `/:productId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goPriceWatch/controllers:ProductTypeController"] = append(beego.GlobalControllerRouter["goPriceWatch/controllers:ProductTypeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
