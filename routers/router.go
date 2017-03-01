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
	beego.Router("/checkin_m", &controllers.CheckinController{}, "get,post:MobileGet")

	// 请假信息
	beego.Router("/leave_details", &controllers.LeaveController{}, "get:PcGet")
	beego.Router("/leave_details_m", &controllers.LeaveController{}, "get:MobileGet")

	beego.Router("/leave_my", &controllers.LeaveController{}, "get:LeaveHistroy")

	beego.Router("/leave_for_leave", &controllers.LeaveController{}, "get,post:AskForLeave")
	beego.Router("/leave_appr_leave", &controllers.LeaveController{}, "get,post:ApproveLeave")
}
