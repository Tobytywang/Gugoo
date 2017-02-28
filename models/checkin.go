package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 打卡信息
type Checkin struct {
	Id     int
	Staff  *Staff `orm:"rel(fk);on_delete(cascade)"` // 用户ID
	Date   string // 打卡日期
	First  int    // 第一次打卡
	Second int    // 第二次打卡
	Third  int    // 第三次打卡
}

const (
	PRE_TIME = 30
)

// 查看所有打卡信息
// 参数： 一个可以容纳这些打卡信息的slice
// 返回： 无
func LoadCheckin(clist *[]*Checkin) {
	o := orm.NewOrm()
	o.QueryTable("checkin").All(clist)
}

// 通过Checkin进行打卡操作
// 参数： 指向用户的指针
// 返回： 状态标记，错误信息
// 标记： 1.不到打卡时间（0:00到8:30，9:00到1:00，1:30到11.59
//       2.打卡成功（上午：8:30到9:00，中午：1:00到1:30，下午：6:00到6:30）
//       3.已经打过卡
//       -1.打卡失败或程序出错
func Check(userid string) (flag int, err error) {
	o := orm.NewOrm()
	fsh, _ := beego.AppConfig.Int("FirstStartHour")
	fsm, _ := beego.AppConfig.Int("FirstStartMinute")
	fs, err := hm2m(fsh, fsm)
	if err != nil {
		return -1, err
	}
	ssh, _ := beego.AppConfig.Int("SecondStartHour")
	ssm, _ := beego.AppConfig.Int("SecondStartMinute")
	ss, err := hm2m(ssh, ssm)
	if err != nil {
		return -1, err
	}
	tsh, _ := beego.AppConfig.Int("ThirdStartHour")
	tsm, _ := beego.AppConfig.Int("ThirdStartMinute")
	ts, err := hm2m(tsh, tsm)
	if err != nil {
		return -1, err
	}
	//nowhour := time.Now().Hour()
	//nowminute := time.Now().Minute()
	now := time.Now().Hour()*60 + time.Now().Minute()
	beego.Debug(time.Now().Hour()*60 + time.Now().Minute())

	checkin := new(Checkin)
	staff, err := StaffByUserId(userid)
	if err != nil {
		return -1, err
	}
	checkin.Staff = staff
	checkin.Date = time.Now().Format("20060102")

	var ch Checkin
	beego.Debug(staff.Id)
	err = o.QueryTable("checkin").Filter("staff_id", staff.Id).Filter("date", time.Now().Format("20060102")).One(&ch)
	// beego.Debug(ch)
	// beego.Debug(err)
	// beego.Debug(reflect.TypeOf(orm.ErrNoRows))
	if err == orm.ErrNoRows {
		beego.Debug("没有查到数据")
		if now <= fs && now >= (fs-PRE_TIME) {
			checkin.First = 1
		} else if now <= ss && now >= (ss-PRE_TIME) {
			checkin.Second = 1
		} else if now <= ts && ts >= (ts-PRE_TIME) {
			checkin.Third = 1
		} else {
			beego.Debug("不在打卡时间")
			return 1, err
		}
		if _, err := o.Insert(checkin); err != nil {
			beego.Debug(err)
			return -1, err
		} else {
			beego.Debug(checkin)
			return 2, nil
		}
	} else {
		beego.Debug(ch)
		ch.Staff = staff
		if now <= fs && now >= (fs-PRE_TIME) {
			beego.Debug(now)
			return 3, nil
		} else if now <= ss && now >= (ss-PRE_TIME) {
			if ch.Second == 1 {
				beego.Debug(now)
				return 3, nil
			} else {
				ch.Second = 1
				if _, err := o.Update(&ch); err != nil {
					beego.Debug(err)
					return -1, err
				} else {
					return 2, nil
				}
			}
		} else if now <= ts && ts >= (ts-PRE_TIME) {
			if ch.Third == 1 {
				beego.Debug(now)
				beego.Debug()
				return 3, nil
			} else {
				ch.Third = 1
				if _, err := o.Update(&ch); err != nil {
					beego.Debug(err)
					return -1, err
				} else {
					return 2, nil
				}
			}
		} else {
			return 1, nil
		}
	}
	return
}

func hm2m(hour int, minute int) (int, error) {
	if (hour >= 0) && (hour <= 23) {
		if (minute >= 0) && (minute <= 59) {
			return (hour*60 + minute), nil
		} else {
			return -1, errors.New("错误的分钟数")
		}
	} else {
		return -1, errors.New("错误的小时数")
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
