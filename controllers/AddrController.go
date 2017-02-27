package controllers

import (
	"Gugoo/models"

	"github.com/astaxie/beego"
)

type AddrController struct {
	beego.Controller
}

func (c *AddrController) Prepare() {

}

func (c *AddrController) Get() {

	// 获得所有的人员
	slist := make([]*models.Staff, 0)
	models.LoadStaff(&slist)
	c.Data["Staff"] = slist

	c.TplName = "pc/addrlist.html"
}
