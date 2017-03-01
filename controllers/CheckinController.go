package controllers

import (
	"Gugoo/models"

	"github.com/astaxie/beego"
)

type CheckinController struct {
	beego.Controller
}

func (c *CheckinController) Prepare() {

}

func (c *CheckinController) MobileGet() {
	if c.Ctx.Input.IsAjax() {
		clist, _ := models.LoadCheckinByTime(c.GetString("year"), c.GetString("month"))
		c.Data["Checkin"] = clist
		beego.Debug("处理ajax")
		c.ServeJSON()
		// c.StopRun()
	} else {
		beego.Debug("不是AJAX")
		clist, _ := models.LoadCheckin()
		c.Data["Checkin"] = clist
	}

	c.TplName = "mobile/checkin.html"
}

func (c *CheckinController) PcGet() {

	// 获得所有的人员
	clist, _ := models.LoadCheckin()
	c.Data["Checkin"] = clist

	c.TplName = "pc/checkin.html"
}
