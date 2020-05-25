package controllers

import (
	"fmt"
	"goPriceWatch/database"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PriceController struct {
	beego.Controller
}

// @Title get
// @Description get object
// @Param id string path true "productid"
// @Param id string path true "days"
// @Success 200 {object} database.Price
// @Failure 403 body is empty
// @router /:productId/:days [get]
func (controller *PriceController) GetForProduct() {
	fmt.Println("Call Id Query")
	var productId = controller.Ctx.Input.Param(":productId")
	var days = controller.Ctx.Input.Param(":days")
	var dayStr, err = strconv.Atoi(days)
	now := time.Now()
	subtractedDate := now.AddDate(0, 0, -dayStr)

	orm.Debug = true
	o := orm.NewOrm()

	var price []*database.Price
	num, err := o.QueryTable("price").Filter("productid", productId).Filter("date__gt", subtractedDate.Format("2006-01-02")).OrderBy("date").All(&price)

	fmt.Printf("Returned Rows Num: %d, %s", num, err)

	data := getJsonFor(price)

	controller.Data["json"] = data
	controller.ServeJSON()
}
