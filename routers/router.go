package routers

import (
	"Gugoo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Login")

	// 打卡信息
	beego.Router("/checkin", &controllers.CheckinController{}, "get:PcGet")
	beego.Router("/checkin_history", &controllers.CheckinController{}, "get,post:MobileGet")

	// 请假信息
	beego.Router("/leave", &controllers.LeaveController{}, "get:PcGet")
	beego.Router("/leave_detail", &controllers.LeaveController{}, "get,post:LeaveDetailApprove")
	beego.Router("/leave_for_leave", &controllers.LeaveController{}, "get,post:AskForLeave")
	beego.Router("/leave_history", &controllers.LeaveController{}, "get:LeaveHistory")
	beego.Router("/leave_to_appr", &controllers.LeaveController{}, "get:WaitApproveList")
	beego.Router("/leave_appr_history", &controllers.LeaveController{}, "get,post:ApproveHistoryList")

}
