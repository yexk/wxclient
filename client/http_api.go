package client

import (
	"encoding/json"
	"io/ioutil"

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
	SEND_PIC_API APIPath = "/api/sendpic"
	// 发消息
	SEND_TXT_MSG_API APIPath = "/api/sendtxtmsg"
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
	GET_PERSONAL_INFO_API APIPath = "/api/get_personal_info"
)

// 获取群成员昵称
func (wx *WxClient) ApiGetMembernick(wxid, roomid string) (nick *Nick, err error) {
	msg := &MSG{
		ID:       GetID(),
		Type:     CHATROOM_MEMBER_NICK,
		Wxid:     wxid,
		Roomid:   roomid,
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
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
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, _e)
	json.Unmarshal([]byte(_e.Content), nick)
	return
}

// 群里发at消息
func (wx *WxClient) ApiSendAtMsg(wxid, content, roomid string) (event *Event, err error) {
	nickname, _ := wx.ApiGetMembernick(wxid, roomid)
	msg := &MSG{
		ID:       GetID(),
		Type:     AT_MSG,
		Wxid:     wxid,
		Roomid:   roomid,
		Content:  content,
		Nickname: nickname.Nick,
		Ext:      "null",
	}
	event = &Event{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(SEND_AT_MSG_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, event)
	return
}

// 发送文本消息
func (wx *WxClient) ApiSendTxtMsg(wxid, content string) (event *Event, err error) {
	msg := &MSG{
		ID:       GetID(),
		Type:     TXT_MSG,
		Wxid:     wxid,
		Roomid:   "null",
		Content:  content,
		Nickname: "null",
		Ext:      "null",
	}
	event = &Event{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(SEND_TXT_MSG_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, event)
	return
}

// 发送图片消息
func (wx *WxClient) ApiSendPicMsg(wxid, content string) (event *Event, err error) {
	msg := &MSG{
		ID:       GetID(),
		Type:     PIC_MSG,
		Wxid:     wxid,
		Roomid:   "null",
		Content:  content,
		Nickname: "null",
		Ext:      "null",
	}
	event = &Event{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(SEND_PIC_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, event)
	return
}

// 获取好友列表
func (wx *WxClient) ApiGetUserList() (event *EventUserList, err error) {
	msg := &MSG{
		ID:       GetID(),
		Type:     USER_LIST,
		Wxid:     "null",
		Roomid:   "null",
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
	}
	event = &EventUserList{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(GET_CONTACTLIST_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, event)
	return
}

// 获取群列表
func (wx *WxClient) ApiGetChatRoomList() (member *EventChatrooMmember, err error) {
	msg := &MSG{
		ID:       GetID(),
		Type:     CHATROOM_MEMBER,
		Wxid:     "null",
		Roomid:   "null",
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
	}
	member = &EventChatrooMmember{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(GET_CHATROOM_MEMBER_LIST_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, member)
	return
}

// 获取个人资料
func (wx *WxClient) ApiGetPersonalInfo() (info *PersonalInfo, err error) {
	msg := &MSG{
		ID:       GetID(),
		Type:     PERSONAL_INFO,
		Wxid:     "null",
		Roomid:   "null",
		Content:  "null",
		Nickname: "null",
		Ext:      "null",
	}
	_e := &content{}
	info = &PersonalInfo{}

	requestBody := GetRequestMsg(msg)
	url := wx.GetHttpUrl(GET_PERSONAL_INFO_API)
	resp, err := common.HttpPost(url, requestBody, nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, _e)
	json.Unmarshal([]byte(_e.Content), info)
	return
}
