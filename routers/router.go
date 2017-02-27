package routers

import (
	"Gugoo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// 打卡信息
	beego.Router("/checkin", &controllers.CheckinController{}, "get:PcGet")
	beego.Router("/checkin_m", &controllers.CheckinController{}, "get:MobileGet")

	// 请假信息
	beego.Router("/leave", &controllers.LeaveController{}, "get:PcGet")
	beego.Router("/leave_m", &controllers.LeaveController{}, "get:MobileGet")
	beego.Router("/leave_asf", &controllers.LeaveController{}, "get,post:AskForLeave")

	// 通讯录
	beego.Router("/addr", &controllers.AddrController{})

}
