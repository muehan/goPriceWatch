package controllers

import (
	"fmt"
	"goPriceWatch/database"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductByTypeController struct {
	beego.Controller
}

// @Title get
// @Description get object
// @Param id string path true "product type id"
// @Success 200 {object} database.Product
// @Failure 403 body is empty
// @router /:productTypeId [get]
func (this *ProductByTypeController) GetProductsByType() {
	fmt.Println("Call Id Query")
	var productTypeId = this.Ctx.Input.Param(":productTypeId")

	orm.Debug = true
	o := orm.NewOrm()

	var products []*database.Product
	num, err := o.QueryTable("product").Filter("Producttypeid", productTypeId).All(&products)
	fmt.Println(num)
	if err != nil {
		this.Abort("404")
	}

	data := getJsonFor(products)

	this.Data["json"] = data
	this.ServeJSON()
}
