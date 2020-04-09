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
	var dates []string
	var data []float64

	orm.Debug = true
	o := orm.NewOrm()

	num1, err := o.QueryTable("product").Filter("Id", id).All(&product)
	fmt.Println("Returned Rows Num: %s, %s", num1, err)
	
	num2, err := o.QueryTable("price").Filter("productid", id).All(&prices)
	fmt.Println("Returned Rows Num: %s, %s", num2, err)

	var max float64
	err = o.Raw("SELECT max(price) FROM price where productid = ?", id).QueryRow(&max)

	var min float64
	err = o.Raw("SELECT min(price) FROM price where productid = ?", id).QueryRow(&min)

	for _, price := range prices {
		date := price.Date.Format("2006-01-02")
		// fmt.Println(date);
		dates = append(dates, date)
		data = append(data, price.Price)
	}

	jsonDates := getJsonFor(dates)
	jsonData := getJsonFor(data)

	c.Data["product"] = product
	c.Data["prices"] = prices
	c.Data["max"] = max
	c.Data["min"] = min
	c.Data["dates"] = jsonDates
	c.Data["data"] = jsonData

	c.TplName = "product/details.html"
	c.Layout = "_layout.html"
}