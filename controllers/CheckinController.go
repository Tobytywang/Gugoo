package controllers

import (
	"Gugoo/models"
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego"
)

type CheckinController struct {
	beego.Controller
}

func (c *CheckinController) Prepare() {

}

func (c *CheckinController) MobileGet() {
	if c.Ctx.Input.IsPost() {
		Clist, _ := models.LoadCheckinByTimeAndUserId("123", c.GetString("year"), c.GetString("month"))
		c.Data["Checkin"] = Clist
		beego.Debug("是POST方法")
	} else {
		beego.Debug(reflect.TypeOf(c.Ctx.Input.IsPost()))
		beego.Debug("是GET方法")
		beego.Debug(fmt.Sprintf("%d", time.Now().Year()), fmt.Sprintf("%02d", time.Now().Month()))
		Clist, _ := models.LoadCheckinByTimeAndUserId("123", fmt.Sprintf("%d", time.Now().Year()), fmt.Sprintf("%02d", time.Now().Month()))
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
