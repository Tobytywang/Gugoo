package controllers

import (
	"Gugoo/models"
	"strconv"
	"time"

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

// 时间格式化函数
// 参数2017-03-03T14:02
func (c *BaseController) GetTime(ti string) time.Time {
	beego.Debug(ti)
	year, _ := strconv.Atoi(c.Substr(ti, 0, 4))
	month, _ := strconv.Atoi(c.Substr(ti, 5, 7))
	day, _ := strconv.Atoi(c.Substr(ti, 8, 10))
	hour, _ := strconv.Atoi(c.Substr(ti, 11, 13))
	minute, _ := strconv.Atoi(c.Substr(ti, 14, 16))
	beego.Debug(year, "-", month, "-", day, "-", hour, "-", minute)
	return time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.Local)
}

func (c *BaseController) Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)
	if start < 0 || start > length {
		return "0"
	}
	if end < 0 || end > length {
		return "0"
	}
	// beego.Debug(string(rs[start:end]))
	return string(rs[start:end])
}
