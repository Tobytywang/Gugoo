package main

import (
	"Gugoo/models"
	_ "Gugoo/routers"
	"fmt"
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

	fmt.Printf("现在的时间是：%d点%d分\n", time.Now().Hour(), time.Now().Minute())
	if n, err := models.Check(); err != nil {
		switch n {
		case 1:
			fmt.Printf("第%d次打卡成功。\n", n)
		case 2:
			fmt.Printf("第%d次打卡成功。\n", n)
		case 3:
			fmt.Printf("第%d次打卡成功。\n", n)
		}
	} else {
		fmt.Printf("打卡失败！\n")
	}

	beego.Run()
}
