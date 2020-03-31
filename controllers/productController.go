package controllers

import (
	"goPriceWatch/database"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) GetProducts() {
	orm.Debug = true
	o := orm.NewOrm()

	var products []*database.Product
	num, err := o.QueryTable("product").All(&products)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	c.Data["products"] = products
	
	c.TplName = "product/index.html"
	c.Layout = "_layout.html"
}

// @router /product/:id
func (c *ProductController) Details() {
	var id = c.Ctx.Input.Param(":id")
	fmt.Println("details: ",id)
	c.Data["id"] = id

	c.TplName = "product/details.html"
	c.Layout = "_layout.html"
}