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

// 通过Checkin进行打卡操作
// 参数为用户Id,时间由函数自动生成
func Check(userid int) (flag int, err error) {
	fsh := beego.AppConfig.Int("FirstStartHour")
	fsm := beego.AppConfig.Int("FirstStartMinute")
	fs := fsh*60 + fsm
	feh := beego.AppConfig.Int("FirstEndHour")
	fem := beego.AppConfig.Int("FirstEndMinute")
	fe := feh*60 + fem
	ssh := beego.AppConfig.Int("SecondStartHour")
	ssm := beego.AppConfig.Int("SecondStartMinute")
	ss := ssh*60 + ssm
	seh := beego.AppConfig.Int("SecondEndHour")
	sem := beego.AppConfig.Int("SecondEndMinute")
	se := seh*60 + sem
	tsh := beego.AppConfig.Int("ThirdStartHour")
	tsm := beego.AppConfig.Int("ThirdStartMinute")
	ts := tsh*60 + tsm
	teh := beego.AppConfig.Int("ThirdEndHour")
	tem := beego.AppConfig.Int("ThirdEndMinute")
	te := teh*60 + tem

	nowhour := time.Now().Hour()
	nowminute := time.Now().Minute()
	now := nowhour*60 + nowminute

	if now <= fs {
		return 1, null
	} else if now >= fe && now <= ss {
		return 2, null
	} else if now >= se && now <= ts {
		return 3, null
	} else {
		return -1, error.New("打卡失败！")
	}
}

// 注册模型
func init() {
	orm.RegisterModel(new(Checkin))
}
