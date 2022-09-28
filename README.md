# wx client dll by golang 

## 安装 
```
go get github.com/yexk/wxclient 
```

## 使用案例
```golang
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

// 好友列表
func RevcUserList(wx *client.WxClient, event *client.EventUserList) {
	for _, v := range event.Content {
		fmt.Printf("v.Name: %v, ", v.Name)
		fmt.Printf("v.Wxcode: %v\n", v.Wxcode)
	}
}

```

## 声明

1. 请大家在国家法律、法规和腾讯相关原则下使用
2. 不对任何下载和使用者的任何行为负责
3. 无任何后门、木马，也不获取、存储任何信息
4. 持续升级
5. 永久免费

## 致谢
[cixingguangming55555](https://github.com/cixingguangming55555)