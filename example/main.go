package main

import (
	"fmt"

	"github.com/yexk/wxclient/client"
)

func main() {
	// 启动监听的地址
	wx := client.NewConnect("127.0.0.1:5555")

	// 注册监听事件
	wx.RegisterHandlers(client.RecvTxtMsgHandler(Revc))
	wx.RegisterHandlers(client.UserListHandler(RevcUserList))

	wx.Start()
}

func Revc(wx *client.WxClient, event *client.Event) {
	// 判断是不是群
	if !wx.BoolGroup(event.Wxid) {
		wx.SendTxtMsg(event.Wxid, "群聊：我是自动回复！")
	} else {
		wx.SendTxtMsg(event.Wxid, "私聊：我是自动回复！")
	}
}

func RevcUserList(wx *client.WxClient, event *client.EventUserList) {
	for _, v := range event.Content {
		fmt.Printf("v.Name: %v, v.Wxcode: %v\n", v.Name, v.Wxcode)
	}
}
