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
	user1.Id = 1
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
	if n, err := models.Check("123"); err != nil {
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
	beego.Debug(models.StaffByUserId("123"))

	leave := new(models.Leave)
	leave.Staff = user1
	leave.ApprovedBy = user1
	leave.DateAsk = time.Now()
	leave.DateOk = time.Now()
	leave.DateStart = time.Now()
	leave.DateEnd = time.Now()
	if err := models.AskLeave(leave); err != nil {
		beego.Debug(err)
	}

	// models.ApproveLeave(user1, leave)
	// 读取员工
	slist := make([]*models.Staff, 0)
	models.LoadStaff(&slist)
	for i := 0; i < len(slist); i++ {
		beego.Debug(slist[i])
	}
	// 读取
	clist := make([]*models.Checkin, 0)
	models.LoadCheckin(&clist)
	for i := 0; i < len(clist); i++ {
		beego.Debug(clist[i])
	}
	// 读取
	llist := make([]*models.Leave, 0)
	models.LoadLeave(&llist)
	for i := 0; i < len(llist); i++ {
		beego.Debug(llist[i])
	}
	beego.Run()
}
