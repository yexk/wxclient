package client

import "github.com/google/uuid"

const HEART_BEAT = 5005               // 心跳
const RECV_TXT_MSG = 1                // 收到的文本消息
const RECV_PIC_MSG = 3                // 收到的图片消息
const RECV_TXT_REFERENCE_MSG = 49     // 引用消息
const USER_LIST = 5000                // 用户和群组列表
const GET_USER_LIST_SUCCSESS = 5001   // 获取成功列表
const GET_USER_LIST_FAIL = 5002       // 获取失败列表
const TXT_MSG = 555                   // 文本消息
const PIC_MSG = 500                   // 图片消息
const AT_MSG = 550                    // 艾特消息
const CHATROOM_MEMBER = 5010          // 群组成员
const CHATROOM_MEMBER_NICK = 5020     // 群组成员别名
const PERSONAL_INFO = 6500            // 个人资料
const DEBUG_SWITCH = 6000             // 调试模式, 默认关
const PERSONAL_DETAIL = 6550          // 个人资料详情
const DESTROY_ALL = 9999              //
const NEW_FRIEND_REQUEST = 37         // 微信好友请求消息
const AGREE_TO_FRIEND_REQUEST = 10000 // 拍一拍
const ATTATCH_FILE = 5003             // 附件

func GetID() string {
	return uuid.New().String()
}
