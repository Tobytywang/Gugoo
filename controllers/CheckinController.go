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
	c.TplName = "mobile/checkin.html"
}

func (c *CheckinController) PcGet() {

	// 获得所有的人员
	clist := make([]*models.Checkin, 0)
	models.LoadCheckin(&clist)
	c.Data["Checkin"] = clist

	c.TplName = "pc/checkin.html"
}
