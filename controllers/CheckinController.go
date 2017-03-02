package controllers

import (
	"Gugoo/models"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type CheckinController struct {
	BaseController
}

func (c *CheckinController) CheckinPrepare() {
	beego.Debug("in CheckinPrepare")
}

func (c *CheckinController) MobileGet() {
	if c.Ctx.Input.IsPost() {
		time := strings.Split(c.GetString("time"), " ")
		beego.Debug(time)
		Clist, _ := models.LoadCheckinByTimeAndUserId(c.UserId, time[0], time[1])
		c.Data["Checkin"] = Clist
		c.Data["ThisYear"] = time[0]
		c.Data["ThisMonth"] = time[1]
		beego.Debug("是POST方法")
	} else {
		beego.Debug(reflect.TypeOf(c.Ctx.Input.IsPost()))
		beego.Debug("是GET方法")

		year := fmt.Sprintf("%d", time.Now().Year())
		month := fmt.Sprintf("%02d", time.Now().Month())
		beego.Debug(year, month)
		Clist, _ := models.LoadCheckinByTimeAndUserId(c.UserId, year, month)
		c.Data["Checkin"] = Clist
		c.Data["ThisYear"] = year
		c.Data["ThisMonth"] = month

	}

	c.TplName = "mobile/checkin.html"
}

func (c *CheckinController) PcGet() {

	// 获得所有的人员
	clist, _ := models.LoadCheckin()
	c.Data["Checkin"] = clist

	c.TplName = "pc/checkin.html"
}
