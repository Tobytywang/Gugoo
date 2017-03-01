package wechat

import (
	"fmt"

	"Gugoo/models"

	"github.com/astaxie/beego/orm"
	"github.com/chanxuehong/wechat/corp/addresslist"
	"github.com/gogather/com/log"
)

type UserPhoneList struct {
	Name        string
	PhoneNumber string
}

//获取简洁通讯录，姓名及电话
func GetAddressList() string {
	addressClient := (*addresslist.Client)(corpClient)
	userList, err := addressClient.UserList(DepartmentId, false, 0)
	if err != nil {
		fmt.Println("获取用户列表错误！")
	}
	userPhoneList := ""
	for _, user := range userList {
		userPhoneList += "\n" + user.Name + ": " + user.Mobile
	}

	return userPhoneList
}

//从微信服务器获取员工信息，并在本地数据库更新
func UpdateStaffInfo() error {
	addressClient := (*addresslist.Client)(corpClient)
	userList, err := addressClient.UserList(DepartmentId, false, 0)
	log.Println(userList)
	if err != nil {
		fmt.Println("获取用户列表错误！")
		return err
	}
	for _, user := range userList {
		staff := &models.Staff{}
		staff.UserId = user.Id
		staff.WeixinId = user.WeixinId
		staff.Name = user.Name
		//目前我们只用一个部门
		staff.Department = user.Department[0]
		staff.Position = user.Position
		staff.Mobile = user.Mobile
		staff.Email = user.Email
		log.Println(staff)

		err := orm.NewOrm().QueryTable("staff").Filter("user_id", staff.UserId).One(staff)
		if err != nil {
			_, err := models.SaveStaff(staff)
			if err != nil {
				return err
			}
		} else {
			err := models.StaffUpdate(staff)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
