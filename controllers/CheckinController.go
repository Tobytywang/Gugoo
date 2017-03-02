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
		Clist, _ := models.LoadCheckinByTimeAndUserId(c.UserId, time[0], time[1])
		c.Data["Checkin"] = Clist
		beego.Debug("是POST方法")
	} else {
		beego.Debug(reflect.TypeOf(c.Ctx.Input.IsPost()))
		beego.Debug("是GET方法")
		beego.Debug(fmt.Sprintf("%d", time.Now().Year()), fmt.Sprintf("%02d", time.Now().Month()))
		Clist, _ := models.LoadCheckinByTimeAndUserId(c.UserId, fmt.Sprintf("%d", time.Now().Year()), fmt.Sprintf("%02d", time.Now().Month()))
		c.Data["Checkin"] = Clist

	}

	c.TplName = "mobile/checkin.html"
}

func (c *CheckinController) PcGet() {

	// 获得所有的人员
	clist, _ := models.LoadCheckin()
	c.Data["Checkin"] = clist

	c.TplName = "pc/checkin.html"
}
