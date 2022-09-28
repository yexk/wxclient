package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/yexk/wxclient/common"
)

type RequetBody struct {
	Para MSG `json:"para"`
}

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
	GET_CHARROOM_MEMBER_LIST_API APIPath = "/api/get_charroom_member_list"
	// 获取个人信息
	GET_PERSONAL_INFO_API APIPath = "/api/get_personal_info"
)

// Tip: API接口与ws重复，暂时不考虑实现，需要的参考下面实现即可
func (wx *WxClient) ApiGetMembernick(wxid, roomid string) (nick *Nick, err error) {
	msg := &MSG{
		ID:       GetID(),
		Type:     CHATROOM_MEMBER_NICK,
		Wxid:     wxid,
		Roomid:   roomid,
		Content:  "",
		Nickname: "null",
		Ext:      "null",
	}
	type e struct {
		Content string `json:"content"`
	}
	_e := &e{}
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
