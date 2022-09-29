package client

import "log"

// 实现对应方法即可实现监听
func eventAll(wx *WxClient, event *Event) {
	log.Printf("channel: %d, handler: %v \n", event.Type, event)
}

func event(wx *WxClient, event *Event) {}

func eventUserList(wx *WxClient, event *EventUserList) {
	log.Printf("userlist: %d, handler: %v \n", event.Type, event)
}

func eventChatRoom(wx *WxClient, event *EventChatrooMmember) {
	log.Printf("chatroom: %d, handler: %v \n", event.Type, event)
}

func eventChatRoomNick(wx *WxClient, event *EventChatrooMmemberNick) {
	log.Printf("nickname: %d, handler: %+v \n", event.Type, event)
}
