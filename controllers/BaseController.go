package controllers

import (
	"Gugoo/models"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	Staff    *models.Staff
	UserId   string //微信的UserId
	UserName string
}

func (c *BaseController) Prepare() {

}
