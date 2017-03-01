package controllers

import (
	"Gugoo/models"

	"Gugoo/wechat"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Staff    *models.Staff
	UserId   string //微信的UserId
	UserName string //对应的姓名
}

//定义接口
type LeavePrepare interface {
	LeavePrepare()
}
type CheckinPrepare interface {
	CheckinPrepare()
}

func (c *BaseController) Prepare() {

	//c.CheckLogin()

	beego.Debug(c.UserId, c.UserName)

	switch app := c.AppController.(type) {
	case LeavePrepare:
		app.LeavePrepare()
	case CheckinPrepare:
		app.CheckinPrepare()
	}
}

func (c *BaseController) CheckLogin() {
	isLogin := c.GetSession("UserId")
	beego.Debug(isLogin)
	if isLogin == nil || isLogin.(string) == "" {
		//beego.Debug("跳转login")
		//c.Redirect(beego.URLFor("LoginController.Login"), 302)
		beego.Debug("第一次登陆")
		requestURI := c.Ctx.Request.RequestURI
		beego.Debug(requestURI)

		//微信企业号登陆入口
		code := c.GetString("code")

		if len(code) > 0 {
			userId, deviceId, err := wechat.GetUserInfo(code)
			if userId != "" && deviceId != "" && err == nil {
				c.SetSession("UserId", userId)
				beego.Debug(userId, deviceId)
				return
			}
			beego.Error("未通过微信验证！")
			return
		}
		redirectURL := wechat.GetAuthCodeURL(wechat.Domain + "/login")
		//redirectURL := c.URLFor("LoginController.Login")
		beego.Debug("redirectURL", redirectURL)
		//wechat.SendText("67", redirectURL)
		c.Redirect(redirectURL, 302)
		return
	}
	uid := isLogin.(string)
	staff, err := models.StaffByUserId(uid)
	if err != nil {
		beego.Debug("不存在该用户！")
	} else {
		c.Staff = staff
		c.UserId = uid
		c.UserName = staff.Name
	}
}
