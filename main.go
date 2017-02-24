package main

import (
	"Gugoo/models"
	_ "Gugoo/routers"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.RegisterDB()
}

func main() {
	beego.Run()
}
