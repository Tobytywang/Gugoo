package controllers

import (
	"Gugoo/models"
	"reflect"

	"github.com/astaxie/beego"
)

type CheckinController struct {
	beego.Controller
}

func (c *CheckinController) Prepare() {

}

func (c *CheckinController) MobileGet() {
	if c.Ctx.Input.IsPost() {
		Clist, _ := models.LoadCheckinByTime(c.GetString("year"), c.GetString("month"))
		c.Data["Checkin"] = Clist
		beego.Debug("是POST方法")
	} else {
		beego.Debug(reflect.TypeOf(c.Ctx.Input.IsPost()))
		beego.Debug("是GET方法")
		Clist, _ := models.LoadCheckin()
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
