package controllers

import (
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
	c.TplName = "pc/checkin.html"
}
