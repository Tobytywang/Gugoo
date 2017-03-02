package wechat

import (
	"github.com/astaxie/beego"
	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/menu"
	"github.com/chanxuehong/wechat/corp/message/request"
	"github.com/chanxuehong/wechat/util"
)

func Wechat() {

	aesKey, err := util.AESKeyDecode(EncodingAESKey)
	if err != nil {
		panic(err)
	}
	messageServeMux := corp.NewMessageServeMux()
	messageServeMux.MessageHandleFunc(request.MsgTypeText, TextMessageHandler) // 注册文本处理 Handler
	//messageServeMux.MessageHandleFunc(request.MsgTypeLocation, LocationMessageHandler) // 注册位置处理 Handler，发送位置消息也可打卡
	messageServeMux.EventHandleFunc(request.EventTypeLocation, LocationEventHandler) // 注册位置事件处理 Handler
	messageServeMux.EventHandleFunc(request.EventTypeSubscribe, SubscribeEventHandler)
	messageServeMux.EventHandleFunc(request.EventTypeUnsubscribe, UnSubscribeEventHandler)
	messageServeMux.EventHandleFunc(menu.EventTypeClick, ClickEventHandler)

	// 下面函数的几个参数设置成你自己的参数: corpId, agentId, token
	agentServer := corp.NewDefaultAgentServer(CorpId, HelperAgentId /* agentId */, HelperToken, aesKey, messageServeMux)

	agentServerFrontend := corp.NewAgentServerFrontend(agentServer, corp.ErrorHandlerFunc(ErrorHandler), nil)

	// 如果你在微信后台设置的回调地址是
	//   http://xxx.yyy.zzz/agent
	// 那么可以这么注册 http.Handler
	beego.Handler(HelperUrl, agentServerFrontend)
	//http.Handle(HelperUrl, agentServerFrontend)
	//http.ListenAndServe(HelperPort, nil)
}
