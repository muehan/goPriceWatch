package controllers

import (
	"github.com/astaxie/beego"
)

type ProductController struct {
	beego.Controller
}

func (c *ProductController) Get() {
	c.Data["products"] = []string{"some link", "someotherlink", "morelinks"}
	
	c.TplName = "product/index.html"
	c.Layout = "_layout.html"
}
