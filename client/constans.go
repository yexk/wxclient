package client

import "github.com/google/uuid"

const (
	// 心跳
	HEART_BEAT = 5005
	// 收到的文本消息
	RECV_TXT_MSG = 1
	// 收到的图片消息
	RECV_PIC_MSG = 3
	// 引用消息
	RECV_TXT_REFERENCE_MSG = 49
	// 用户和群组列表
	USER_LIST = 5000
	// 获取成功列表
	GET_USER_LIST_SUCCSESS = 5001
	// 获取失败列表
	GET_USER_LIST_FAIL = 5002
	// 文本消息
	TXT_MSG = 555
	// 图片消息
	PIC_MSG = 500
	// 艾特消息
	AT_MSG = 550
	// 群组成员
	CHATROOM_MEMBER = 5010
	// 群组成员别名
	CHATROOM_MEMBER_NICK = 5020
	// 个人资料
	PERSONAL_INFO = 6500
	// 调试模式, 默认关
	DEBUG_SWITCH = 6000
	// 个人资料详情
	PERSONAL_DETAIL = 6550
	//
	DESTROY_ALL = 9999
	// 微信好友请求消息
	NEW_FRIEND_REQUEST = 37
	// 添加好友，拍一拍
	AGREE_TO_FRIEND_REQUEST = 10000
	// 附件
	ATTATCH_FILE = 5003
)

// uuid
func GetID() string {
	return uuid.New().String()
}
