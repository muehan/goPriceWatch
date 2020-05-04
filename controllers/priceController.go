package controllers

import (
	"goPriceWatch/database"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type PriceController struct {
	beego.Controller
}

// @Title get
// @Description get object
// @Param id string path true "productid"
// @Success 200 {object} database.Price
// @Failure 403 body is empty
// @router /:productId [get]
func (this *PriceController) GetForProduct() {
	fmt.Println("Call Id Query")
	var productId = this.Ctx.Input.Param(":productId")

	orm.Debug = true
	o := orm.NewOrm()

	var price []*database.Price
	num, err := o.QueryTable("price").Filter("productid", productId).OrderBy("date").All(&price)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)

	data := getJsonFor(price)
	
	this.Data["json"] = data
    this.ServeJSON()
}
