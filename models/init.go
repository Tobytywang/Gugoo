package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_DRIVER = "mysql"
)

// 构造链接数据库的字符串
var _DB_CONNECT_STR string = beego.AppConfig.String("mysqluser") +
	":" + beego.AppConfig.String("mysqlpass") + "@/" +
	beego.AppConfig.String("mysqldb") + "?charset=utf8" + "&loc=Local"

// 注册模型，建表
func RegisterDB() {
	orm.RegisterDriver(_DB_DRIVER, orm.DRMySQL)
	orm.RegisterDataBase("default", _DB_DRIVER, _DB_CONNECT_STR)
	orm.RunSyncdb("default", false, true)
}
