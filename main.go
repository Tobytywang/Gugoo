package main

import (
	"Gugoo/models"
	_ "Gugoo/routers"

	"Gugoo/wechat"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/chanxuehong/wechat/corp/message/request"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.RegisterDB()
	orm.RunSyncdb("default", false, true)
	wechat.UpdateStaffInfo()
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = false
	//beego.SetLevel(beego.LevelWarning)
	// 自动建表
	//orm.RunSyncdb("default", false, true)

	//wechat.CreateMenu() //修改菜单时用一次就好
	//wechat.PrintMenu()

	wechat.LocationMap = make(map[string]request.LocationEvent)
	go wechat.Wechat()

	beego.Run()
}
