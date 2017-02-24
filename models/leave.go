package models

import (
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

// 注册模型
func init() {
	orm.RegisterModel(new(Leave))
}
