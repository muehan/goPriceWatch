package controllers

import (
	"fmt"
	"goPriceWatch/database"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ProductTypeController struct {
	beego.Controller
}

// @Title get
// @Description get object
// @Success 200 {object} database.Producttype
// @Failure 403 body is empty
// @router / [get]
func (this *ProductTypeController) Get() {
	orm.Debug = true
	o := orm.NewOrm()

	var producttypes []*database.Producttype
	num, err := o.QueryTable("producttype").All(&producttypes)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)

	this.Data["json"] = producttypes
	this.ServeJSON()
}
