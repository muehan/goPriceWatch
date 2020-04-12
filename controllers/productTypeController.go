package controllers

import (
	"goPriceWatch/database"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
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

	data := getJsonFor(producttypes)
	
	this.Data["json"] = data
    this.ServeJSON()
}
