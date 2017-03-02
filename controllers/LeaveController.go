package controllers

import (
	"Gugoo/models"
	"Gugoo/wechat"
	"time"

	"github.com/astaxie/beego"
)

type LeaveController struct {
	BaseController
}

func (c *LeaveController) LeavePrepare() {
	beego.Debug("in LeavePrepare")
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
			leave.Staff = c.Staff
			// 批准人
			leave.ApprovedBy, _ = models.StaffByUserId(c.GetString("whoappr"))
			leave.DateAsk = time.Now()
			leave.DateOk = leave.DateAsk //将批准时间设置和申请时间相同，表示待批准
			leave.DateStart = c.GetTime(c.GetString("start"))
			leave.DateEnd = c.GetTime(c.GetString("end"))
			leave.Reason = c.GetString("reason")
			//---------------------------------------------
			err := models.LeaveAdd(&leave)
			beego.Debug("leaveAdd ", err)
			if err != nil {
				flash.Error("申请失败，请重试！")
				flash.Store(&c.Controller)
				c.Redirect("/leave_for_leave", 302)
			} else {
				beego.Debug("申请成功！")
				flash.Notice("申请成功！")
				flash.Store(&c.Controller)
				wechat.SendText(leave.ApprovedBy.UserId, leave.Staff.Name+"向你发起了请假申请，请在待我审批菜单里查看详情！")
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

// 查看假条详情，并处理
func (c *LeaveController) LeaveDetailApprove() {
	leaveid, err := c.GetInt("leaveid")
	beego.Debug(leaveid, err)
	leave, err := models.LeaveGetById(leaveid)
	if err != nil {
		beego.Debug(err)
		beego.Warn("查找假条失败！")
	}
	if c.Ctx.Input.IsPost() {
		isPass := c.GetString("PassOrNot")
		if len(isPass) > 0 {
			msg := "你的请假申请审批结果是："
			switch isPass {
			case "yes":
				leave.ApprovedState = 1
				msg += "通过\n"
			case "no":
				leave.ApprovedState = -1
				msg += "未通过\n"
			}
			msg += "详细情况请在请假记录里查看"
			wechat.SendText(leave.Staff.UserId, msg)
			models.LeaveUpdate(leave, "ApprovedState")
		}
		c.Redirect("/leave_to_appr", 302) //转到待我审批列表

	} else if c.Ctx.Input.IsGet() {
		op := c.GetString("op")
		c.Data["Leave"] = leave
		switch op {
		case "viewleave":
			c.Data["op"] = "viewleave"
		case "approve":
			c.Data["op"] = "approve"
		case "viewapprove":
			c.Data["op"] = "viewapprove"
		}

		c.TplName = "mobile/detail.html"
	}

}

// 查看我的请假历史
func (c *LeaveController) LeaveHistory() {
	ls, err := models.LeaveListGetByAskStaffId(c.UserId)
	if err != nil {
		beego.Debug(err)
	}
	c.Data["table_cate"] = "LeaveHistory"
	c.Data["List"] = ls

	c.TplName = "mobile/approvalRecord.html"
}

// 查看待我审批表单
func (c *LeaveController) WaitApproveList() {

	ls, err := models.WaitingLeaveListGetByApprovedStaffId(c.UserId, false)
	if err != nil {
		beego.Debug(err)
	}
	c.Data["table_cate"] = "WaitApproveList"
	c.Data["List"] = ls

	c.TplName = "mobile/approvalRecord.html"
}

// 查看审批记录表单
func (c *LeaveController) ApproveHistoryList() {

	ls, err := models.WaitingLeaveListGetByApprovedStaffId(c.UserId, true)
	if err != nil {
		beego.Debug(err)
	}
	c.Data["table_cate"] = "ApproveHistoryList"
	c.Data["List"] = ls

	c.TplName = "mobile/approvalRecord.html"
}
