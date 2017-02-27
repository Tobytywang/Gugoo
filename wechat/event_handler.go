package wechat

import (
	"fmt"
	"log"
	"math"
	"net/http"

	//"time"

	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/menu"
	"github.com/chanxuehong/wechat/corp/message/request"
	"github.com/chanxuehong/wechat/corp/message/response"
	//"Gugoo/models"
	"time"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err.Error())
}

//根据经纬度求距离
func EarthDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6378137)
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

// 位置事件的 Handler,将每个人的位置信息记录下来，目前用全局变量或缓存，然后根据30s内的位置信息判断打卡
func LocationEventHandler(w http.ResponseWriter, r *corp.Request) {
	location := request.GetLocationEvent(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	LocationMap[location.FromUserName] = *location
}

// 订阅事件的 Handler
func SubscribeEventHandler(w http.ResponseWriter, r *corp.Request) {
	subscribe := request.GetSubscribeEvent(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	fmt.Println(subscribe)
	resp := response.NewText(subscribe.FromUserName, subscribe.ToUserName, subscribe.CreateTime, "欢迎关注Gugoo微信企业号～～～")
	corp.WriteResponse(w, r, resp)
}

// 取消订阅事件的 Handler
func UnSubscribeEventHandler(w http.ResponseWriter, r *corp.Request) {
	//用户取消关注后，需要做些什么？？？
	unsubscribe := request.GetUnsubscribeEvent(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	resp := response.NewText(unsubscribe.FromUserName, unsubscribe.ToUserName, unsubscribe.CreateTime, "别走呀～～～")
	fmt.Println(unsubscribe, resp)
}

// 点击事件的handler
func ClickEventHandler(w http.ResponseWriter, r *corp.Request) {
	click := menu.GetClickEvent(r.MixedMsg)
	loc := LocationMap[click.FromUserName]
	distance := EarthDistance(LocationX, LocationY, loc.Latitude, loc.Longitude)
	switch click.EventKey {
	case "1": //打卡
		msg := ""
		//state,err := models.Check(click.FromUserName)
		switch {
		//缓存的位置信息超过30秒，不算
		case time.Now().Unix()-loc.CreateTime > 30:
			msg = "打卡失败\n\n尚未获取到你当前的位置信息，请检查是否已允许提供位置信息并重试！"
			break
		//case ://已打过卡
		//	msg =  "已打过卡，请勿重复操作！"
		//	break
		//将打卡距离设置为50米
		case distance > 50:
			msg = "打卡失败\n\n时间：" + time.Unix(loc.CreateTime, 0).Format("2006-01-02 15:04:05") + "\n距离：" + fmt.Sprint(distance) + "米\n\n距离工作室太远，请进入工作室再重试！"
			break
		//case distance <= 50://在工作室附近，后台数据库insert失败
		//	msg = "打卡失败\n在工作室附近，后台数据库insert失败，请联系后台开发人员（卢琦或王天宇）"
		//	break
		case distance <= 50: //
			msg = "打卡成功\n\n时间：" + time.Unix(loc.CreateTime, 0).Format("2006-01-02 15:04:05") + "\n距离：" + fmt.Sprint(distance) + "米"
			break
		}

		resp := response.NewText(click.FromUserName, click.ToUserName, click.CreateTime, msg)
		corp.WriteResponse(w, r, resp)

	case "2": //从数据库搜索今日打卡状态，分为上午、下午、晚上
		resp := response.NewText(click.FromUserName, click.ToUserName, click.CreateTime, "今日打卡情况\n\n"+"上午：已打卡\n"+"下午：已打卡\n"+"晚上：未打卡"+"\n" /*+GetAuthCodeURL()*/)
		corp.WriteResponse(w, r, resp)

	case "3": //发送手机通讯录
		userPhoneList := GetAddressList()
		resp := response.NewText(click.FromUserName, click.ToUserName, click.CreateTime, "简版通讯录\n"+userPhoneList)
		corp.WriteResponse(w, r, resp)
	}

}