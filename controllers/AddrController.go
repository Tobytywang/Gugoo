package controllers

import (
	"github.com/astaxie/beego"
)

type AddrController struct {
	beego.Controller
}

func (c *AddrController) Prepare() {

}

func (c *AddrController) Get() {

	c.TplName = "pc/addrlist.html"
}
