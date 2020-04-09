package controllers

import (
	"goPriceWatch/database"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type ProductTypeApiController struct {
	beego.Controller
}

func (this *ProductTypeApiController) Get() {
	orm.Debug = true
	o := orm.NewOrm()

	var producttypes []*database.Producttype
	num, err := o.QueryTable("producttype").All(&producttypes)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	data := getJsonFor(producttypes)
	
	this.Data["json"] = data
    this.ServeJSON()
}
