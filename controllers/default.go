package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
	c.Layout = "_layout.html"
}

func getJsonFor(object interface{}) string {
	// fmt.Println(object)
	var jsonData []byte
	jsonData, err := json.Marshal(object)
	if err != nil {
    	fmt.Println(err)
	}

	return string(jsonData)
}