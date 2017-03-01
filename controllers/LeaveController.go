package controllers

import (
	"Gugoo/models"

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

//请假入口
func (c *LeaveController) AskForLeave() {

	leave := new(models.Leave)

	err := models.LeaveAdd(leave)
	if err != nil {

	}
	beego.Debug(err)

	c.TplName = "mobile/askforleave.html"
}

//待审批入口，审批人在这里审批，然后更新假条信息
func (c *LeaveController) WaitApprove() {
	c.TplName = "mobile/askforleave.html"
}

//已审批入口，审批人查看审批记录
func (c *LeaveController) ApprovedHistoryLeave() {
	c.TplName = "mobile/askforleave.html"
}

//查看请假记录入口
func (c *LeaveController) LeaveHistroy() {
	c.TplName = "mobile/askforleave.html"
}
