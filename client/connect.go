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
		type e struct {
			Type int32 `json:"type"`
		}
		ev := &e{}
		err = json.Unmarshal(msg[:], ev)
		if err != nil {
			log.Printf("new Connect json.Unmarshal err: %v", err)
			continue
		}
		event := &Event{}
		if ev.Type == RECV_TXT_REFERENCE_MSG {
			_event := &EventReferenceConent{}
			json.Unmarshal(msg[:], _event)
			// 做一个解析转换
			event = &Event{
				ID:       _event.ID,
				Wxid:     _event.Content.ID1,
				Content:  _event.Content.Content,
				Type:     _event.Type,
				Receiver: _event.Receiver,
				Sender:   _event.Sender,
				Srvid:    _event.Srvid,
				Status:   _event.Status,
				Time:     _event.Time,
			}
			if wx.BoolGroup(_event.Content.ID1) {
				event.Wxid = _event.Content.ID1
				event.ID1 = _event.Content.ID2
			}
		} else if ev.Type == USER_LIST || ev.Type == GET_USER_LIST_FAIL || ev.Type == GET_USER_LIST_SUCCSESS {
			_event := &EventUserList{}
			json.Unmarshal(msg[:], _event)
			wx.handlerUserList(_event)
			continue
		} else if ev.Type == CHATROOM_MEMBER_NICK {
			_event := &EventChatrooMmemberNick{}
			// json字符串特殊处理
			_e := &content{}
			_nick := &Nick{}
			json.Unmarshal(msg[:], _e)
			json.Unmarshal([]byte(_e.Content), _nick)
			// ----处理结束
			json.Unmarshal(msg[:], _event)
			// 赋值
			_event.Content = *_nick
			wx.handlerChatrooMmemberNick(_event)
			continue
		} else if ev.Type == CHATROOM_MEMBER {
			_event := &EventChatrooMmember{}
			json.Unmarshal(msg[:], _event)
			wx.handlerChatrooMmember(_event)
			continue
		} else if ev.Type == AGREE_TO_FRIEND_REQUEST {
			_event := &EventPaiyipai{}
			json.Unmarshal(msg[:], _event)
			// 数据转换
			event = &Event{
				ID:       _event.ID,
				Wxid:     _event.Content.ID1,
				Content:  _event.Content.Content,
				Type:     _event.Type,
				Receiver: _event.Receiver,
				Sender:   _event.Sender,
				Srvid:    _event.Srvid,
				Status:   _event.Status,
				Time:     _event.Time,
			}
		} else if ev.Type == RECV_PIC_MSG {
			_event := &EventPicConent{}
			json.Unmarshal(msg[:], _event)
			// 做一个解析转换
			event = &Event{
				ID:       _event.ID,
				Wxid:     _event.Content.ID1,
				Content:  _event.Content.Content,
				Type:     _event.Type,
				Receiver: _event.Receiver,
				Sender:   _event.Sender,
				Srvid:    _event.Srvid,
				Status:   _event.Status,
				Time:     _event.Time,
			}
			if wx.BoolGroup(_event.Content.ID1) {
				event.Wxid = _event.Content.ID1
				event.ID1 = _event.Content.ID2
			}
		} else {
			json.Unmarshal(msg[:], event)
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

func (wx *WxClient) handlerUserList(event *EventUserList) {
	switch event.Type {
	case USER_LIST:
		go Handlers.UserList(wx, event)
	case GET_USER_LIST_SUCCSESS:
		go Handlers.UserList(wx, event)
	case GET_USER_LIST_FAIL:
		go Handlers.UserList(wx, event)
	}
}

func (wx *WxClient) handlerChatrooMmember(event *EventChatrooMmember) {
	go Handlers.ChatRoomMember(wx, event)
}

func (wx *WxClient) handlerChatrooMmemberNick(event *EventChatrooMmemberNick) {
	go Handlers.ChatRoomMemberNick(wx, event)
}
