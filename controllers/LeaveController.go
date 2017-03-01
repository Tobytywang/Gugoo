package controllers

import (
	"Gugoo/models"
	"time"

	"github.com/astaxie/beego"
)

type LeaveController struct {
	BaseController
}

func (c *LeaveController) LeavePrepare() {
	beego.Warn("in LeavePrepare")
}

// 移动端和PC端查看请假信息
func (c *LeaveController) MobileGet() {
	c.TplName = "mobile/approvalRecord.html"
}
func (c *LeaveController) PcGet() {
	c.TplName = "pc/leave.html"
}

// 发起请假请求
func (c *LeaveController) AskForLeave() {
	flash := beego.NewFlash()
	if c.Ctx.Input.IsPost() {
		if c.GetString("reason") == "" || c.GetString("start") == "" || c.GetString("end") == "" {
			flash.Error("提交的信息不能为空！")
			flash.Store(&c.Controller)
			c.Redirect("/leave_for_leave", 302)
		} else {
			leave := models.Leave{}
			//----------------------------------------------
			// 请假人
			leave.Staff, _ = models.StaffByUserId("123")
			// 批准人
			leave.ApprovedBy, _ = models.StaffByUserId("123")
			leave.DateAsk = time.Now()
			leave.DateOk = leave.DateAsk
			leave.DateStart = c.GetTime(c.GetString("start"))
			leave.DateEnd = c.GetTime(c.GetString("end"))
			//---------------------------------------------
			err := models.LeaveAdd(&leave)
			if err != nil {
				flash.Error("提交的信息有误，请核对后再次提交！")
				flash.Store(&c.Controller)
				c.Redirect("/leave_for_leave", 302)
			}
		}
	} else {
		flash = beego.ReadFromRequest(&c.Controller)
		if _, ok := flash.Data["error"]; ok {
			c.Data["Error"] = true
		}
	}
	c.TplName = "mobile/forLeave.html"
}

// 处理请假请求
func (c *LeaveController) ApproveLeave() {
	leaveid := c.GetString("leaveid")

	c.TplName = "mobile/detail.html"
}

// 查看我的请假历史
func (c *LeaveController) LeaveHistroy() {
	c.TplName = "mobile/approvalRecord.html"
}
