package controllers

import (
	"fmt"
	"goPriceWatch/database"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductByDateController struct {
	beego.Controller
}

// @Title getbydate
// @Description get object
// @Param date string path true "date"
// @Success 200 {object} database.Product
// @Failure 403 body is empty
// @router /:date [get]
func (this *ProductByDateController) GetByDate() {
	fmt.Println("Call date Query")
	var date = this.Ctx.Input.Param(":date")

	orm.Debug = true
	o := orm.NewOrm()

	var products []*database.Product
	num, err := o.QueryTable("product").Filter("date", date).All(&products)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)

	this.Data["json"] = products
	this.ServeJSON()
}
