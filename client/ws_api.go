package client

// 发文本消息
func (wx *WxClient) SendTxtMsg(wxid, content string) (i int, err error) {
	i, err = wx.Send(&MSG{
		Wxid:    wxid,
		Content: content,
	})
	return
}

// 发图片消息
func (wx *WxClient) SendPicMsg(wxid, content string) (i int, err error) {
	i, err = wx.Send(&MSG{
		Wxid:    wxid,
		Content: content,
	})
	return
}

// 发送艾特消息
func (wx *WxClient) SendAtMsg(wxid, content, roomid string) (i int, err error) {
	// nickname, _ := wx.ApiGetMembernick(wxid, roomid)
	i, err = wx.Send(&MSG{
		Wxid:    wxid,
		Content: content,
	})
	return
}

// 获取个人信息
func (wx *WxClient) GetPersonalInfo(wxid string) (i int, err error) {
	i, err = wx.Send(&MSG{
		Wxid:    wxid,
		Content: "null",
	})
	return
}

// 获取详细资料
func (wx *WxClient) GetPersonalDetail(wxid string) (i int, err error) {
	i, err = wx.Send(&MSG{
		Wxid:    wxid,
		Content: "null",
	})
	return
}

// 获取群成员昵称
func (wx *WxClient) GetGroupUserNick(wxid, roomid string) (i int, err error) {
	i, err = wx.Send(&MSG{
		Wxid:    wxid,
		Content: "",
	})
	return
}

// 获取群列表
func (wx *WxClient) GetChatRoomMemberList() (i int, err error) {
	i, err = wx.Send(&MSG{
		Wxid:    "null",
		Content: "null",
	})
	return
}

// 获取微信好友列表
func (wx *WxClient) GetUserList() (i int, err error) {
	i, err = wx.Send(&MSG{
		Wxid:    "null",
		Content: "null",
	})
	return
}
