package client

// 处理事件函数
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
	AddFriendRequest   AddFriendRequestHandler
	TxtMsg             TxtMsgHandler
	PicMsg             PicMsgHandler
	AtMsg              AtMsgHandler
}

type (
	// ReadyHandler 可以处理所有的事件
	ReadyHandler func(wx *WxClient, event *Event)

	// RecvTxtMsgHandler 获取接收文本消息事件
	RecvTxtMsgHandler func(wx *WxClient, event *Event)

	// RecvAtMsgHandler 获取接收At消息事件
	RecvAtMsgHandler func(wx *WxClient, event *Event)

	// RecvPicMsgHandler 获取接收图片消息事件
	RecvPicMsgHandler func(wx *WxClient, event *Event)

	// UserListHandler 获取用户列表消息事件
	UserListHandler func(wx *WxClient, event *EventUserList)

	// ChatRoomMemberHandler 获取群成员事件
	ChatRoomMemberHandler func(wx *WxClient, event *EventChatrooMmember)

	// ChatRoomMemberNickHandler 获取成员具体昵称事件
	ChatRoomMemberNickHandler func(wx *WxClient, event *EventChatrooMmemberNick)

	// PersonalInfoHandler 获取个人资料事件
	PersonalInfoHandler func(wx *WxClient, event *Event)

	// PersonalDetailHandler 获取个人详细资料事件
	PersonalDetailHandler func(wx *WxClient, event *Event)

	// AddFriendRequestHandler 获取添加好友事件
	AddFriendRequestHandler func(wx *WxClient, event *Event)

	// AgreeToFriendHandler 同意好友成功事件
	AgreeToFriendHandler func(wx *WxClient, event *Event)

	// TxtMsgHandler 文本消息事件
	TxtMsgHandler func(wx *WxClient, event *Event)

	// PicMsgHandler 图片消息事件
	PicMsgHandler func(wx *WxClient, event *Event)

	// AtMsgHandler At消息事件
	AtMsgHandler func(wx *WxClient, event *Event)
)

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
		case AddFriendRequestHandler:
			Handlers.AddFriendRequest = handle
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
