package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 打卡信息
type Checkin struct {
	Id     int
	Staff  *Staff    `orm:"rel(fk);on_delete(cascade)"` // 用户ID
	Date   time.Time `orm:"type(date)"`                 // 打卡日期
	First  int       // 第一次打卡
	Second int       // 第二次打卡
	Third  int       // 第三次打卡
}

// 查看所有打卡信息
// 参数： 一个可以容纳这些打卡信息的slice
// 返回： 无
func LoadCheckin(clist *[]*Checkin) {
	o := orm.NewOrm()
	o.QueryTable("checkin").All(clist)
}

// 通过Checkin进行打卡操作
// 参数： 指向用户的指针
// 返回： 打卡记录，错误信息
func Check(userid string) (flag int, err error) {
	o := orm.NewOrm()
	fsh, _ := beego.AppConfig.Int("FirstStartHour")
	fsm, _ := beego.AppConfig.Int("FirstStartMinute")
	fs := fsh*60 + fsm
	feh, _ := beego.AppConfig.Int("FirstEndHour")
	fem, _ := beego.AppConfig.Int("FirstEndMinute")
	fe := feh*60 + fem

	ssh, _ := beego.AppConfig.Int("SecondStartHour")
	ssm, _ := beego.AppConfig.Int("SecondStartMinute")
	ss := ssh*60 + ssm
	seh, _ := beego.AppConfig.Int("SecondEndHour")
	sem, _ := beego.AppConfig.Int("SecondEndMinute")
	se := seh*60 + sem

	tsh, _ := beego.AppConfig.Int("ThirdStartHour")
	tsm, _ := beego.AppConfig.Int("ThirdStartMinute")
	ts := tsh*60 + tsm
	//teh, _ := beego.AppConfig.Int("ThirdEndHour")
	//tem, _ := beego.AppConfig.Int("ThirdEndMinute")
	//te := teh*60 + tem

	nowhour := time.Now().Hour()
	nowminute := time.Now().Minute()
	now := nowhour*60 + nowminute

	checkin := new(Checkin)
	if checkin.Staff, err = StaffByUserId(userid); err != nil {
		return -1, err
	}
	checkin.Date = time.Now()

	beego.Debug(checkin.Staff)
	beego.Debug(*checkin)

	if now <= fs { //00:00-09:00
		checkin.First = 1
		if _, err := o.Insert(checkin); err != nil {
			return -1, err
		} else {
			return 0, nil
		}
	} else if now >= fe && now <= ss { //11:30-13:30
		checkin.Second = 1
		if _, err := o.Insert(checkin); err != nil {
			return -1, err
		} else {
			return 0, nil
		}
	} else if now >= se && now <= ts { //17:00-18:00
		checkin.Third = 1
		if _, err := o.Insert(checkin); err != nil {
			return -1, err
		} else {
			return 0, nil
		}
	} else { //不在打卡时间内
		return -1, err
	}
}

func (c *Checkin) TableUnique() [][]string {
	return [][]string{
		[]string{"Staff", "Date"},
	}
}

// 注册模型
func init() {
	orm.RegisterModel(new(Checkin))
}
