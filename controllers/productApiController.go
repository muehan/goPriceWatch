package controllers

import (
	"goPriceWatch/database"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type ProductApiController struct {
	beego.Controller
}

func (this *ProductApiController) Get() {
	orm.Debug = true
	o := orm.NewOrm()

	var product []*database.Product
	num, err := o.QueryTable("product").All(&product)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	data := getJsonFor(product)
	
	this.Data["json"] = data
    this.ServeJSON()
}

func (this *ProductApiController) GetDetails() {
	var id = this.Ctx.Input.Param(":id")

	orm.Debug = true
	o := orm.NewOrm()

	var product []*database.Product
	num, err := o.QueryTable("product").Filter("Id", id).All(&product)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	fmt.Println(product)

	data := getJsonFor(product)
	
	this.Data["json"] = data
    this.ServeJSON()
}
