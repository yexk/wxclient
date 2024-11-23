package client

import (
	"encoding/json"
	"io"

	"github.com/yexk/wxclient/common"
)

// 组装请求body
type requetBody struct {
	Para MSG `json:"para"`
}

// Api路径
type APIPath string

const (
	// at消息接口
	SEND_AT_MSG_API APIPath = "/api/sendatmsg"
	// 图片消息接口
	SEND_PIC_API APIPath = "/api/send_pic"
	// 发消息
	SEND_TXT_MSG_API APIPath = "/api/send_txt"
	// 发附件
	SEND_ATTATCH_API APIPath = "/api/sendattatch"
	// 获取别名
	GET_MEMBERNICK_API APIPath = "/api/getmembernick"
	// 获取成员ID
	GET_MEMBERID_API APIPath = "/api/getmemberid"
	// 获取通讯录
	GET_CONTACTLIST_API APIPath = "/api/getcontactlist"
	// 获取成员列表
	GET_CHATROOM_MEMBER_LIST_API APIPath = "/api/get_charroom_member_list"
	// 获取个人信息
	GET_PERSONAL_INFO_API APIPath = "/api/get_self"
)

// 获取群成员昵称
func (wx *WxClient) ApiGetMembernick(wxid, roomid string) (nick *Nick, err error) {
	msg := &MSG{
		Wxid:    wxid,
		Content: "null",
	}
	_e := &content{}
	nick = &Nick{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(GET_MEMBERNICK_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, _e)
	json.Unmarshal([]byte(_e.Content), nick)
	return
}

// 群里发at消息
func (wx *WxClient) ApiSendAtMsg(wxid, content, roomid string) (event *Event, err error) {
	// nickname, _ := wx.ApiGetMembernick(wxid, roomid)
	msg := &MSG{
		Wxid:    wxid,
		Content: content,
	}
	event = &Event{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(SEND_AT_MSG_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, event)
	return
}

// 发送文本消息
func (wx *WxClient) ApiSendTxtMsg(wxid, content string) (event *Event, err error) {
	msg := &MSG{
		Wxid:    wxid,
		Content: content,
	}
	respEvent := &RespEvent{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(SEND_TXT_MSG_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, respEvent)

	json.Unmarshal([]byte(respEvent.Data), event)
	return
}

// 发送图片消息
func (wx *WxClient) ApiSendPicMsg(wxid, content string) (event *Event, err error) {
	msg := &MSG{
		Wxid:    wxid,
		Content: content,
	}
	event = &Event{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(SEND_PIC_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, event)
	return
}

// 获取好友列表
func (wx *WxClient) ApiGetUserList() (event *EventUserList, err error) {
	msg := &MSG{
		Wxid:    "null",
		Content: "null",
	}
	event = &EventUserList{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(GET_CONTACTLIST_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, event)
	return
}

// 获取群列表
func (wx *WxClient) ApiGetChatRoomList() (member *EventChatrooMmember, err error) {
	msg := &MSG{
		Wxid:    "null",
		Content: "null",
	}
	member = &EventChatrooMmember{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(GET_CHATROOM_MEMBER_LIST_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, member)
	return
}

// 获取个人资料
func (wx *WxClient) ApiGetPersonalInfo() (info *PersonalInfo, err error) {
	_e := &RespEvent{}
	info = &PersonalInfo{}

	url := wx.GetHttpUrl(GET_PERSONAL_INFO_API)
	resp, err := common.HttpPost(url, "", nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, _e)
	json.Unmarshal([]byte(_e.Data), info)
	return
}
