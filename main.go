package main

import (
	"Gugoo/models"
	_ "Gugoo/routers"
	"time"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.RegisterDB()
}

func main() {
	user1 := new(models.Staff)
	user1.UserId = "123"
	user1.Name = "HappyLich"
	user1.Department = 3
	user1.Position = "liu"
	user1.Mobile = "13333333333"
	user1.Email = "wang@163.com"
	user1.WeixinId = "455sd"

	if _, err := models.SaveStaff(user1); err != nil {
		beego.Debug("内部错误:")
		beego.Debug(err)
	}
	beego.Debug("现在的时间是：", time.Now().Hour(), "点", time.Now().Minute(), "分\n")
	if n, err := models.Check(1); err != nil {
		beego.Debug(err)
	} else {
		switch n {
		case 1:
			beego.Debug("第", n, "次打卡成功。")
		case 2:
			beego.Debug("第", n, "次打卡成功。")
		case 3:
			beego.Debug("第", n, "次打卡成功。")
		default:
			beego.Debug("n=", n)
		}
	}

	tm := time.Date(2017, 2, 25, 8, 30, 0, 0, time.Local)
	beego.Debug("现在假设有时间：", tm.Hour(), "点", tm.Minute(), "分\n")
	if n, err := models.Check(1); err != nil {
		beego.Debug(err)
	} else {
		switch n {
		case 1:
			beego.Debug("第", n, "次打卡成功。")
		case 2:
			beego.Debug("第", n, "次打卡成功。")
		case 3:
			beego.Debug("第", n, "次打卡成功。")
		default:
			beego.Debug("n=", n)
		}
	}

	beego.Debug("根据ID查找用户\n")
	beego.Debug(models.StaffById(1))
	beego.Run()
}
