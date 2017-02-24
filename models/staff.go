package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 创建成员的参数
type Staff struct {
	Id         int
	UserId     string `json:"userid,omitempty"`     // 必须;  员工UserID. 对应管理端的帐号, 企业内必须唯一. 长度为1~64个字符
	Name       string `json:"name,omitempty"`       // 必须;  成员名称. 长度为1~64个字符
	Department int    `json:"department,omitempty"` // 非必须; 成员所属部门id列表. 注意, 每个部门的直属员工上限为1000个
	Position   string `json:"position,omitempty"`   // 非必须; 职位信息. 长度为0~64个字符
	Mobile     string `json:"mobile,omitempty"`     // 非必须; 手机号码. 企业内必须唯一, mobile/weixinid/email三者不能同时为空
	Email      string `json:"email,omitempty"`      // 非必须; 邮箱. 长度为0~64个字符. 企业内必须唯一
	WeixinId   string `json:"weixinid,omitempty"`   // 非必须; 微信号. 企业内必须唯一. (注意: 是微信号, 不是微信的名字)
}

// 注册模型
func init() {
	orm.RegisterModel(new(Staff))
}
