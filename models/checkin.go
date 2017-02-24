package models

import (
	"time"

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
// function Checkin(userid int) {
//   nowhours := time.Now().Hour() *  + time.Now().Minite()
// }

// 注册模型
func init() {
	orm.RegisterModel(new(Checkin))
}
