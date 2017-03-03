package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 请假信息
type Leave struct {
	Id            int
	Staff         *Staff    `orm:"rel(fk);on_delete(cascade)"form:"whoask"`        // 用户ID
	ApprovedBy    *Staff    `orm:"null;rel(fk);on_delete(set_null)"form:"whoappr"` // 审批人ID
	DateAsk       time.Time `orm:"type(datetime)"`                                 // 申请时间
	DateOk        time.Time `orm:"null;type(datetime)"`                            // 审批时间
	DateStart     time.Time `orm:"type(datetime)"form:"start"`                     // 请假开始时间
	DateEnd       time.Time `orm:"type(datetime)"form:"end"`                       // 请假结束时间
	Reason        string    `form:"reason"`                                        //请假理由
	ApprovedState int       //审批状态，0：待审批，-1：审批不通过，1：审批通过
}

// 查看所有请假信息
// 参数： 一个可以容纳这些请假信息的slice
// 返回： 无
func LoadLeave() (list []Leave) {
	o := orm.NewOrm()
	o.QueryTable("leaves").RelatedSel().All(&list)
	return list
}

// 发起请假申请，添加假条，前台页面表单都不能为空，注意过滤
// 同时存在在数据库中的ApprovedState标志为0的申请，每个用户只能有一个
// 参数： 请假结构Leave
// 返回： 报错信息
func LeaveAdd(leave *Leave) error {
	o := orm.NewOrm()
	//var lv Leave
	//o.QueryTable("leaves").Filter("staff_id", leave.Staff.Id).Filter("approved_state", 0).One(&lv)
	//if lv.Id != 0 {
	//	return errors.New("请等待当前请假申请批准")
	//}
	_, err := o.Insert(leave)
	return err
}

//点击某条记录，进入详情页面，需要根据Id查询假条信息
func LeaveGetById(lid int) (*Leave, error) {
	l := new(Leave)
	err := orm.NewOrm().QueryTable("leaves").Filter("id", lid).RelatedSel().One(l)
	return l, err
}

// 根据ID获得请假信息
// func LoadLeaveById(id int)(llist []Leave, err error){
// 	o := orm.NewOrm()
// 	if id <= 0 {
// 		return nil, errors.New("错误的ID请求")
// 	} else {
// 		var leave Leave
// 		o.QueryTable("leaves").Filter("id", id).One(&leave)
// 		if
// 	}
// }

//根据申请人userId查询请假记录
func LeaveListGetByAskStaffId(userid string) ([]*Leave, error) {
	ls := make([]*Leave, 0)
	_, err := orm.NewOrm().QueryTable("leaves").Filter("Staff__UserId", userid).OrderBy("-id").RelatedSel().All(&ls)
	return ls, err
}

//根据审批人userId查询假条,state（false：待审批，true：已审批）
func WaitingLeaveListGetByApprovedStaffId(userid string, state bool) (ls []*Leave, err error) {
	ls = make([]*Leave, 0)
	switch state {
	case true:
		_, err = orm.NewOrm().QueryTable("leaves").Filter("ApprovedBy__UserId", userid).Filter("ApprovedState__in", 1, -1).OrderBy("-id").RelatedSel().All(&ls)
	case false:
		_, err = orm.NewOrm().QueryTable("leaves").Filter("ApprovedBy__UserId", userid).Filter("ApprovedState", 0).OrderBy("-id").RelatedSel().All(&ls)
	}
	return
}

//审批人审批后需要更新假条信息
func LeaveUpdate(l *Leave, fields ...string) error {
	_, err := orm.NewOrm().Update(l, fields...)
	return err
}

// 审批
// 要求有权限认证机制
// 这个可以在Model里写，可以在Controller里写
// 参数：审批人，假条，审批结果
//func ApproveLeave(appr *Staff, leave *Leave, result bool) error {
//	o := orm.NewOrm()
//	if appr == leave.ApprovedBy {
//		leave.ApprovedState = 1
//		if _, err := o.Update(leave); err != nil {
//			return err
//		} else {
//			return nil
//		}
//	} else {
//		return errors.New("没有权限")
//	}
//}

// 自定义表名
func (u *Leave) TableName() string {
	return "leaves"
}

// 注册模型
func init() {
	orm.RegisterModel(new(Leave))
}
