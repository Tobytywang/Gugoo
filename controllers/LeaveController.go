package controllers

import (
	"github.com/astaxie/beego"
)

type LeaveController struct {
	beego.Controller
}

func (c *LeaveController) Prepare() {

}

func (c *LeaveController) MobileGet() {
	c.TplName = "mobile/leave.html"
}

func (c *LeaveController) PcGet() {
	c.TplName = "pc/leave.html"
}

func (c *LeaveController) AskForLeave() {
	c.TplName = "mobile/askforleave.html"
}
