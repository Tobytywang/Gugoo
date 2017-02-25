package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 请假信息
type Leave struct {
	Id         int
	Staff      *Staff    `orm:"rel(fk);on_delete(cascade)"`       // 用户ID
	ApprovedBy *Staff    `orm:"null;rel(fk);on_delete(set_null)"` // 批准人ID
	DateAsk    time.Time `orm:"type(datetime)"`                   // 申请时间
	DateOk     time.Time `orm:"type(datetime)"`                   // 批准时间
	DateStart  time.Time `orm:"type(datetime)"`                   // 请假开始时间
	DateEnd    time.Time `orm:"type(datetime)"`                   // 请假结束时间
	IsApproved int
}

// 查看所有请假信息
// 参数： 一个可以容纳这些请假信息的slice
// 返回： 无
func LoadLeave(llist *[]*Leave) {
	o := orm.NewOrm()
	o.QueryTable("leaves").All(llist)
}

// 发起请假申请函数
// 同时存在在数据库中的IsApproved标志为0的申请，每个用户只能有一个
// 参数： 请假结构Leave
// 返回： 报错信息
func AskLeave(leave *Leave) error {
	o := orm.NewOrm()
	var lv Leave
	o.QueryTable("leaves").Filter("staff_id", leave.Staff.Id).Filter("is_approved", 0).One(&lv)
	if lv.Id != 0 {
		return errors.New("请等待当前请假申请批准")
	} else {
		if _, err := o.Insert(leave); err != nil {
			return err
		} else {
			return nil
		}
	}
}

// 批准请假
// 要求有权限认证机制
// 这个可以在Model里写，可以在Controller里写
// 参数：
func ApproveLeave(appr *Staff, leave *Leave) error {
	o := orm.NewOrm()
	if appr == leave.ApprovedBy {
		leave.IsApproved = 1
		if _, err := o.Update(leave); err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return errors.New("没有权限")
	}
}

// 自定义表名
func (u *Leave) TableName() string {
	return "leaves"
}

// 注册模型
func init() {
	orm.RegisterModel(new(Leave))
}
