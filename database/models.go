package database

import (
    "github.com/astaxie/beego/orm"
    "time"
)

type Producttype struct {
    Id          string    `orm:"pk"`
    typeid      int
    Name        string
}

type Product struct {
    Id                  string    `orm:"pk"`
    producttypeid       string
    productId           string
    productIdAsString   string
    name                string
    fullname            string
    simpleName          string
    // Prices      []*Price  `orm:"reverse(many)"`
}

type Price struct {
    Id          string     `orm:"pk"`
    Price       float64     // dangerous, but no other solution found so far
    Date        time.Time
    Product     *Product   `orm:"column(productid);rel(one)"` // Reverse relationship (optional)
}

func init() {
    // Need to register model in init
    orm.RegisterModel(new(Product), new(Price), new(Producttype))
}
