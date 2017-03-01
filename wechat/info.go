package wechat

import "github.com/chanxuehong/wechat/corp/message/request"

//咕咕科技企业号的各种信息
const (
	Domain = "http://luqi0119.cn:8080"
	CorpId = "wx02c1979e62f43b58"
	Secret = "jQ0wm3zRxEd1dZLgNwui5Zk_SBiuzizXLVTRUKSZXtGFM8xH1s2awmECgwoBiIZP"

	DepartmentId = 1 // 成都咕咕科技有限公司部门Id

	//应用回调模式，AgentID,EncodingAESKey是一一对应的
	//TestAgentId = 15 //测试
	//TestToken   = "test"
	//TestUrl     = "/test"
	//TestPort    = ":8088"

	HelperAgentId = 0 //企业小助手
	HelperToken   = "helper"
	HelperUrl     = "/helper"
	HelperPort    = ":8081"

	//ClockInAgentId = 22 //打卡
	//ClockInToken   = "clockIn"
	//ClockInUrl     = "/clockIn"
	//ClockInPort    = ":8082"
	//
	//AddressBookAgentId = 24 //通讯录
	//AddressBookToken   = "addressBook"
	//AddressBookUrl     = "/addressBook"
	//AddressBookPort    = ":8083"

	EncodingAESKey = "euQsGnC5KXOtkvug4vuHrivg2TFGcCShLOd2mvGoHH1" //全部设一样的

	//工作室位置信息，经纬度
	LocationX = 30.75486
	LocationY = 103.919052
)

var LocationMap map[string]request.LocationEvent
