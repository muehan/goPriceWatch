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

	var product database.Product
	var prices []*database.Price

	orm.Debug = true
	o := orm.NewOrm()

	num1, err := o.QueryTable("product").Filter("Id", id).All(&product)
	fmt.Printf("Returned Rows Num: %s, %s", num1, err)
	fmt.Printf("Found product: %s", product)

	
	num2, err := o.QueryTable("price").Filter("productid", id).All(&prices)
	fmt.Printf("Returned Rows Num: %s, %s", num2, err)

	c.Data["product"] = product
	c.Data["prices"] = prices

	c.TplName = "product/details.html"
	c.Layout = "_layout.html"
}