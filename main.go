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
}

func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	//orm.RunSyncdb("default", false, true)

	//wechat.CreateMenu()
	wechat.PrintMenu()
	wechat.LocationMap = make(map[string]request.LocationEvent)
	//wechat.UpdateStaffInfo()
	go wechat.Wechat()

	//user1 := new(models.Staff)
	//user1.Id = 1
	//user1.UserId = "123"
	//user1.Name = "HappyLich"
	//user1.Department = 3
	//user1.Position = "liu"
	//user1.Mobile = "13333333333"
	//user1.Email = "wang@163.com"
	//user1.WeixinId = "455sd"
	//
	//if _, err := models.SaveStaff(user1); err != nil {
	//	beego.Debug("内部错误:")
	//	beego.Debug(err)
	//}
	//beego.Debug("现在的时间是：", time.Now().Hour(), "点", time.Now().Minute(), "分\n")
	//if n, err := models.Check("123"); err != nil {
	//	beego.Debug(err)
	//} else {
	//	switch n {
	//	case 1:
	//		beego.Debug("第", n, "次打卡成功。")
	//	case 2:
	//		beego.Debug("第", n, "次打卡成功。")
	//	case 3:
	//		beego.Debug("第", n, "次打卡成功。")
	//	default:
	//		beego.Debug("n=", n)
	//	}
	//}
	//
	//beego.Debug("根据ID查找用户\n")
	//beego.Debug(models.StaffByUserId("123"))

	beego.Run()
}
