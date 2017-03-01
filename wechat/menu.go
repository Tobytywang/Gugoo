package wechat

import (
	"fmt"

	"log"

	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/menu"
	"github.com/chanxuehong/wechat/corp/message/send"
	"github.com/chanxuehong/wechat/corp/oauth2"
)

var AccessTokenServer = corp.NewDefaultAccessTokenServer(CorpId, Secret, nil) // 一個應用只能有一個實例
var corpClient = corp.NewClient(AccessTokenServer, nil)

func CreateMenu() {
	var subButtons0 = make([]menu.Button, 3)
	//subButtons0[0].SetAsLocationSelectButton("位置", "1")
	//subButtons0[1].SetAsClickButton("赞一下", "2")
	//subButtons0[2].SetAsViewButton("博客", "http://luqi0119.cn")
	subButtons0[0].SetAsClickButton("打卡", "1")
	subButtons0[1].SetAsClickButton("状态查询", "2")
	subButtons0[2].SetAsViewButton("历史记录", Domain+"/checkin_m")

	//var subButtons1 = make([]menu.Button, 2)
	//subButtons1[0].SetAsClickButton("手机号", "3")
	//subButtons1[1].SetAsViewButton("详细信息", "http://www.baidu.com")

	//subButtons1[0].SetAsPicPhotoOrAlbumButton("PicOrAlbum", "3")
	//subButtons1[1].SetAsPicSysPhotoButton("SysPhoto", "4")
	//subButtons1[2].SetAsPicWeixinButton("PicWeixin", "5")
	var subButtons2 = make([]menu.Button, 4)
	subButtons2[0].SetAsViewButton("请假", Domain+"/leave_asf")
	subButtons2[1].SetAsViewButton("请假记录", Domain+"/leave_m")
	//这个要匹配，只有管理人员才显示这一项
	subButtons2[2].SetAsViewButton("待我审批", Domain+"/leave_m")
	subButtons2[3].SetAsViewButton("审批记录", Domain+"/leave_m")

	//subButtons2[0].SetAsScanCodePushButton("ScanCodePush", "6")
	//subButtons2[1].SetAsScanCodeWaitMsgButton("ScanCodeWaitMsg", "7")
	//
	var mn menu.Menu
	mn.Buttons = make([]menu.Button, 3)
	mn.Buttons[0].SetAsSubMenuButton("打卡栏", subButtons0)
	mn.Buttons[2].SetAsClickButton("通讯录", "3")
	mn.Buttons[1].SetAsSubMenuButton("请假栏", subButtons2)

	menuClient := (*menu.Client)(corpClient)
	if err := menuClient.CreateMenu(HelperAgentId, mn); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}

func PrintMenu() {
	menuClient := (*menu.Client)(corpClient)
	var mn menu.Menu
	var err error
	if mn, err = menuClient.GetMenu(HelperAgentId); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(mn)

}

//验证时通过次url获取code
func GetAuthCodeURL(url string) string {
	return oauth2.AuthCodeURL(CorpId, url, "snsapi_base", "67")
}

//授权网站验证时用到该函数
func GetUserInfo(code string) (string, string, error) {
	//网页获取用户信息
	oauth2Client := (*oauth2.Client)(corpClient)
	userInfo, err := oauth2Client.UserInfo(HelperAgentId, code)
	log.Println("GetUserInfo", err)
	return userInfo.UserId, userInfo.DeviceId, err
}

func SendText(userid, content string) {
	sendClient := (*send.Client)(corpClient)
	msgheader := &send.MessageHeader{
		MsgType: send.MsgTypeText,
		AgentId: HelperAgentId,
		ToUser:  userid,
	}
	text := new(send.Text)
	text.MessageHeader = *msgheader
	text.Text.Content = content

	_, err := sendClient.SendText(text)
	log.Println(text, err)
}
