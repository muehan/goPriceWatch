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

// @Title get
// @Description get object
// @Success 200 {object} database.Product
// @Failure 403 body is empty
// @router / [get]
func (this *ProductController) Get() {
	fmt.Println("Call all Query")

	orm.Debug = true
	o := orm.NewOrm()

	var product []*database.Product
	num, err := o.QueryTable("product").All(&product)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)

	data := getJsonFor(product)
	
	this.Data["json"] = data
    this.ServeJSON()
}

// @Title get
// @Description get object
// @Param id string path true "product id"
// @Success 200 {object} database.Product
// @Failure 403 body is empty
// @router /:productId [get]
func (this *ProductController) GetDetails() {
	fmt.Println("Call Id Query")
	var productId = this.Ctx.Input.Param(":productId")

	orm.Debug = true
	o := orm.NewOrm()

	var product database.Product
	err := o.QueryTable("product").Filter("productidasstring", productId).One(&product)
	if (err != nil) {
		this.Abort("404")
	}

	fmt.Println(product)

	data := getJsonFor(product)
	
	this.Data["json"] = data
    this.ServeJSON()
}

// @Title get
// @Description get object
// @Param search string query true "search string"
// @Success 200 {object} database.Product
// @Failure 403 body is empty
// @router / [get]
func (this *ProductController) Search() {
	fmt.Println("Call Search Query")
	var search = this.Ctx.Input.Param(":search")

	orm.Debug = true
	o := orm.NewOrm()

	var products database.Product
	num, err := o.QueryTable("product").Filter("fullname__contains", search).All(&products)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	data := getJsonFor(products)
	
	this.Data["json"] = data
    this.ServeJSON()
}