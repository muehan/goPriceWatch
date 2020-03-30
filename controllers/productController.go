package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) GetProducts() {
	c.Data["products"] = []string{"some link", "someotherlink", "morelinks"}
	
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