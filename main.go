package main

import (
	"encoding/json"
	"os"
	"fmt"
	_ "goPriceWatch/database"
	_ "goPriceWatch/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type Configuration struct {
    ConnectionString string
}

func init() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
	  fmt.Println("error:", err)
	}
	fmt.Println(configuration.ConnectionString)

	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", 
	"postgres",
	configuration.ConnectionString)

	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
    o.Using("default")

	beego.Run()
}

