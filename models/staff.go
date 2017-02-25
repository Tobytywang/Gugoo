package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 创建成员的参数
type Staff struct {
	Id         int
	UserId     string `orm:"unique"json:"userid,omitempty"`   // 必须;  员工UserID. 对应管理端的帐号, 企业内必须唯一. 长度为1~64个字符
	Name       string `json:"name,omitempty"`                 // 必须;  成员名称. 长度为1~64个字符
	Department int    `json:"department,omitempty"`           // 非必须; 成员所属部门id列表. 注意, 每个部门的直属员工上限为1000个
	Position   string `json:"position,omitempty"`             // 非必须; 职位信息. 长度为0~64个字符
	Mobile     string `orm:"unique"json:"mobile,omitempty"`   // 非必须; 手机号码. 企业内必须唯一, mobile/weixinid/email三者不能同时为空
	Email      string `orm:"unique"json:"email,omitempty"`    // 非必须; 邮箱. 长度为0~64个字符. 企业内必须唯一
	WeixinId   string `orm:"unique"json:"weixinid,omitempty"` // 非必须; 微信号. 企业内必须唯一. (注意: 是微信号, 不是微信的名字)
}

// 读取函数，从数据库里读出所有的成员
func LoadStaff() {

}

// 存储函数，向数据库里储存所有的成员
// 参数： 一个指向staff的指针类型
// 返回： 错误号码和错误信息
func SaveStaff(staff *Staff) (number int, err error) {
	o := orm.NewOrm()
	if _, err := o.Insert(staff); err != nil {
		return -1, err
	} else {
		return 0, nil
	}
}

// 根据ID查找员工
// 参数： 员工id
// 返回： 员工指针，错误信息
func StaffById(id int) (s *Staff, err error) {
	o := orm.NewOrm()
	var staff Staff
	o.QueryTable("staff").Filter("id", id).One(&staff)

	if staff.Id != 0 {
		return &staff, nil
	} else {
		return &staff, errors.New("没有该用户")
	}
}

// 注册模型
func init() {
	orm.RegisterModel(new(Staff))
}
