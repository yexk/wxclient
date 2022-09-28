package client

var Handlers struct {
	Ready              ReadyHandler
	RecvTxtMsg         RecvTxtMsgHandler
	RecvPicMsg         RecvPicMsgHandler
	UserList           UserListHandler
	ChatRoomMember     ChatRoomMemberHandler
	ChatRoomMemberNick ChatRoomMemberNickHandler
	RecvAtMsg          RecvAtMsgHandler
	PersonalInfo       PersonalInfoHandler
	PersonalDetail     PersonalDetailHandler
	AgreeToFriend      AgreeToFriendHandler
	TxtMsg             TxtMsgHandler
	PicMsg             PicMsgHandler
	AtMsg              AtMsgHandler
}

// ReadyHandler 可以处理所有的事件
type ReadyHandler func(wx *WxClient, event *Event)

// RecvTxtMsgHandler 获取接收文本消息事件
type RecvTxtMsgHandler func(wx *WxClient, event *Event)

// RecvAtMsgHandler 获取接收At消息事件
type RecvAtMsgHandler func(wx *WxClient, event *Event)

// RecvPicMsgHandler 获取接收图片消息事件
type RecvPicMsgHandler func(wx *WxClient, event *Event)

// UserListHandler 获取用户列表消息事件
type UserListHandler func(wx *WxClient, event *EventUserList)

// ChatRoomMemberHandler 获取群成员事件
type ChatRoomMemberHandler func(wx *WxClient, event *EventChatrooMmember)

// ChatRoomMemberNickHandler 获取成员具体昵称事件
type ChatRoomMemberNickHandler func(wx *WxClient, event *EventChatrooMmemberNick)

// PersonalInfoHandler 获取个人资料事件
type PersonalInfoHandler func(wx *WxClient, event *Event)

// PersonalDetailHandler 获取个人详细资料事件
type PersonalDetailHandler func(wx *WxClient, event *Event)

// AgreeToFriendHandler 获取添加好友事件
type AgreeToFriendHandler func(wx *WxClient, event *Event)

// TxtMsgHandler 文本消息事件
type TxtMsgHandler func(wx *WxClient, event *Event)

// PicMsgHandler 图片消息事件
type PicMsgHandler func(wx *WxClient, event *Event)

// AtMsgHandler At消息事件
type AtMsgHandler func(wx *WxClient, event *Event)

// RegisterHandlers 注册事件回调
func (wx *WxClient) RegisterHandlers(handlers ...interface{}) {
	for _, h := range handlers {
		switch handle := h.(type) {
		case ReadyHandler:
			Handlers.Ready = handle
		case RecvTxtMsgHandler:
			Handlers.RecvTxtMsg = handle
		case RecvPicMsgHandler:
			Handlers.RecvPicMsg = handle
		case UserListHandler:
			Handlers.UserList = handle
		case ChatRoomMemberHandler:
			Handlers.ChatRoomMember = handle
		case ChatRoomMemberNickHandler:
			Handlers.ChatRoomMemberNick = handle
		case RecvAtMsgHandler:
			Handlers.RecvAtMsg = handle
		case PersonalDetailHandler:
			Handlers.PersonalDetail = handle
		case PersonalInfoHandler:
			Handlers.PersonalInfo = handle
		case AgreeToFriendHandler:
			Handlers.AgreeToFriend = handle
		case TxtMsgHandler:
			Handlers.TxtMsg = handle
		case PicMsgHandler:
			Handlers.PicMsg = handle
		case AtMsgHandler:
			Handlers.AtMsg = handle
		default:
		}
	}
}
