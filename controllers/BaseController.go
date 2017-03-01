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
	beego.Debug("in prepare")
	c.CheckLogin()

	//AppController表示当前子类是哪个Controller
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
		beego.Debug("第一次登陆")
		firstRequestURI := c.Ctx.Request.RequestURI //记住用户从哪个uri进入的，跳到login界面后，方便跳回来
		beego.Debug(firstRequestURI)

		redirectURL := wechat.GetAuthCodeURL(wechat.Domain + "/login?first=" + firstRequestURI)
		//beego.Debug("redirectURL", redirectURL)
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
