package client

// 发文本消息
func (wx *WxClient) SendTxtMsg(wxid, content string) (i int, err error) {
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     TXT_MSG,
		Wxid:     wxid,
		Roomid:   "null",
		Content:  content,
		Nickname: "null",
		Ext:      "null",
	})
	return
}

// 发图片消息
func (wx *WxClient) SendPicMsg(wxid, content string) (i int, err error) {
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     PIC_MSG,
		Wxid:     wxid,
		Roomid:   "null",
		Content:  content,
		Nickname: "null",
		Ext:      "null",
	})
	return
}

// 发送艾特消息
func (wx *WxClient) SendAtMsg(wxid, content, roomid string) (i int, err error) {
	nickname, _ := wx.ApiGetMembernick(wxid, roomid)
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     AT_MSG,
		Wxid:     wxid,
		Roomid:   roomid,
		Content:  content,
		Nickname: nickname.Nick,
		Ext:      "null",
	})
	return
}

// 获取个人信息
func (wx *WxClient) GetPersonalInfo(wxid string) (i int, err error) {
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     PERSONAL_INFO,
		Wxid:     wxid,
		Roomid:   "null",
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
	})
	return
}

// 获取详细资料
func (wx *WxClient) GetPersonalDetail(wxid string) (i int, err error) {
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     PERSONAL_DETAIL,
		Wxid:     wxid,
		Roomid:   "null",
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
	})
	return
}

// 获取群成员昵称
func (wx *WxClient) GetGroupUserNick(wxid, roomid string) (i int, err error) {
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     CHATROOM_MEMBER_NICK,
		Wxid:     wxid,
		Roomid:   roomid,
		Content:  "",
		Nickname: "null",
		Ext:      "null",
	})
	return
}

// 获取群列表
func (wx *WxClient) GetChatRoomMemberList() (i int, err error) {
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     CHATROOM_MEMBER,
		Wxid:     "null",
		Roomid:   "null",
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
	})
	return
}

// 获取微信好友列表
func (wx *WxClient) GetUserList() (i int, err error) {
	i, err = wx.Send(&MSG{
		ID:       GetID(),
		Type:     USER_LIST,
		Wxid:     "null",
		Roomid:   "null",
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
	})
	return
}
