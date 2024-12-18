package client

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type WxClient struct {
	ws *websocket.Conn

	// 请求地址
	wsUrl   string
	httpUrl string

	once     sync.Once
	RecvWxid string
}

// 初始化机器人，以及建立基础连接
func NewConnect(host string) (wx *WxClient) {
	wx = &WxClient{}
	wx.parse(host)

	origin := "http://localhost/"
	ws, err := websocket.Dial(wx.wsUrl, "", origin)
	if err != nil {
		panic(err)
	}
	wx.ws = ws
	log.Println("connect websocket success!")

	wx.once.Do(func() {
		info, err := wx.ApiGetPersonalInfo()
		if err != nil {
			log.Panicf("login failed: %v", err)
			panic(err)
		}
		log.Printf("login %s success!", info.Name)
	})

	// 注册监听事件
	wx.RegisterHandlers(ReadyHandler(eventAll))
	wx.RegisterHandlers(RecvTxtMsgHandler(event))
	wx.RegisterHandlers(RecvPicMsgHandler(event))
	wx.RegisterHandlers(UserListHandler(eventUserList))
	wx.RegisterHandlers(ChatRoomMemberHandler(eventChatRoom))
	wx.RegisterHandlers(ChatRoomMemberNickHandler(eventChatRoomNick))
	wx.RegisterHandlers(RecvAtMsgHandler(event))
	wx.RegisterHandlers(PersonalInfoHandler(event))
	wx.RegisterHandlers(PersonalDetailHandler(event))
	wx.RegisterHandlers(AgreeToFriendHandler(event))
	wx.RegisterHandlers(TxtMsgHandler(event))
	wx.RegisterHandlers(PicMsgHandler(event))
	wx.RegisterHandlers(AtMsgHandler(event))

	return
}

// 启动监听消息
func (wx *WxClient) Start() {
	// 监听消息
	wx.readHandler()
}

// 监听消息
func (wx *WxClient) readHandler() {
	for {
		var err error
		msg := wx.readMessage()
		if msg == nil {
			time.Sleep(2 * time.Second)
			continue
		}
		// 解析特殊类型
		respEvent := &RespEvent{}
		event := &Event{}
		err = json.Unmarshal(msg, respEvent)
		if respEvent.Code == 200 && err == nil {
			json.Unmarshal([]byte(respEvent.Data), event)
		}

		go Handlers.Ready(wx, event)

		wx.handler(event)
	}
}

// 处理事件
func (wx *WxClient) handler(event *Event) {
	switch event.Type {
	case PERSONAL_INFO:
		go Handlers.PersonalInfo(wx, event)
	case PERSONAL_DETAIL:
		go Handlers.PersonalDetail(wx, event)
	case RECV_TXT_MSG:
		go Handlers.RecvTxtMsg(wx, event)
	case RECV_PIC_MSG:
		go Handlers.RecvPicMsg(wx, event)
	case NEW_FRIEND_REQUEST:
		go Handlers.AddFriendRequest(wx, event)
	case AGREE_TO_FRIEND_REQUEST:
		go Handlers.AgreeToFriend(wx, event)
	case TXT_MSG:
		go Handlers.TxtMsg(wx, event)
	case PIC_MSG:
		go Handlers.PicMsg(wx, event)
	case AT_MSG:
		go Handlers.AtMsg(wx, event)
	}
}
