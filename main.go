package main

import (
	_ "Gugoo/models"
	_ "Gugoo/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/gugoo?charset=utf8&loc=Local", 30)
	orm.RunSyncdb("default", true, false)
}

func main() {
	beego.Run()
}
