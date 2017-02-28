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

func GetTodayCheckinStateByUserid(userid string) (string, error) {
	o := orm.NewOrm()
	var ch Checkin
	var state string = "今日打卡情况\n\n"
	staff, err := StaffByUserId(userid)
	if err != nil {
		return "不存在该用户", err
	}
	err = o.QueryTable("checkin").Filter("staff_id", staff.Id).Filter("date", time.Now().Format("20060102")).One(&ch)

	if err == orm.ErrNoRows {
		return "今日打卡情况\n\n上午：未打卡\n下午：未打卡\n晚上：未打卡", nil
	}

	if err != nil {
		return "程序错误！请联系后台开发人员（卢琦或王天宇）", err
	}
	var mp [2]string
	mp[0] = "未打卡"
	mp[1] = "已打卡"
	state += "上午：" + mp[ch.First] + "\n下午：" + mp[ch.Second] + "\n晚上：" + mp[ch.Third]

	return state, nil
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

	//测试用
	//fs += 60
	//ss += 60
	//ts += 60

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
	beego.Debug(now, fs, ss, ts)
	if err == orm.ErrNoRows {
		beego.Debug("没有查到数据")
		if now <= fs+PRE_TIME && now >= (fs-PRE_TIME) {
			checkin.First = 1
			if _, err := o.Insert(checkin); err != nil {
				beego.Debug(err)
				return -1, err
			} else {
				beego.Debug("上午")
				return 2, nil
			}
		} else if now <= ss+PRE_TIME && now >= (ss-PRE_TIME) {
			checkin.Second = 1
			if _, err := o.Insert(checkin); err != nil {
				beego.Debug(err)
				return -1, err
			} else {
				beego.Debug("中午")
				return 2, nil
			}
		} else if now <= ts+PRE_TIME && now >= (ts-PRE_TIME) {
			checkin.Third = 1
		} else {
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
		if now <= fs+PRE_TIME && now >= (fs-PRE_TIME) {
			beego.Debug(now)
			return 3, nil
		} else if now <= ss+PRE_TIME && now >= (ss-PRE_TIME) {
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
		} else if now <= ts+PRE_TIME && now >= (ts-PRE_TIME) {
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

//将8:30转换为8*60+30=510分钟再来比较大小
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
