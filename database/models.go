package database

import (
    "github.com/astaxie/beego/orm"
    "time"
)

type Producttype struct {
    Id          string    `orm:"pk"`
    Typeid      int
    Name        string
}

type Product struct {
    Id                  string    `orm:"pk"`
    Producttypeid       string    `orm:"column(producttypeid)"`
    Productid           string    `orm:"column(productid)"`
    ProductidAsString   string    `orm:"column(productidasstring)"`
    Name                string    `orm:"column(name)"`
    Fullname            string    `orm:"column(fullname)"`
    SimpleName          string    `orm:"column(simpleName)"`
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
