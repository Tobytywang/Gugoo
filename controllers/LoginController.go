package controllers

import (
	"Gugoo/wechat"

	"github.com/astaxie/beego"
)

type LoginController struct {
	//BaseController
	beego.Controller
}

func (c *LoginController) Login() {
	beego.Debug("in login")
	requestURI := c.Ctx.Request.RequestURI
	beego.Debug(requestURI)
	errMsg := "没有权限！"
	//微信企业号登陆入口
	code := c.GetString("code")
	firstRequestURI := c.GetString("first")
	beego.Debug(code)

	if len(code) > 0 {
		userId, deviceId, err := wechat.GetUserInfo(code)
		if userId != "" && deviceId != "" && err == nil {
			c.SetSession("UserId", userId)
			beego.Debug(userId, deviceId)
			beego.Debug(firstRequestURI)
			c.Redirect(firstRequestURI, 302)
			return
		}
		errMsg = "未通过微信验证！"
	}

	c.Data["json"] = &map[string]interface{}{"error": errMsg, "code": code}
	c.ServeJSON()
	c.StopRun()
}
