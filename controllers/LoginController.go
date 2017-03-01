package controllers

import (
	"Gugoo/wechat"

	"github.com/astaxie/beego"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {
	beego.Debug("in login")
	requestURI := c.Ctx.Request.RequestURI
	beego.Debug(requestURI)

	//微信企业号登陆入口
	code := c.GetString("code")
	beego.Debug(code)

	if len(code) > 0 {
		userId, deviceId, err := wechat.GetUserInfo(code)
		if userId != "" && deviceId != "" && err == nil {
			c.SetSession("UserId", userId)
			beego.Debug(userId, deviceId)
			c.Redirect("/checkin_m", 302)
			return
		}
		beego.Error("未通过微信验证！")
		return
	}
	redirectURL := wechat.GetAuthCodeURL(wechat.Domain + "/login")

	beego.Debug(redirectURL)

	c.Redirect(redirectURL, 302)
}
