package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
)

// 判断是否为群
func (wx *WxClient) BoolGroup(wxid string) bool {
	reg, _ := regexp.MatchString("@chatroom", wxid)
	return reg
}

// 获取ws地址
func (wx *WxClient) GetWsUrl() string {
	return wx.wsUrl
}

// 获取http地址 + path
func (wx *WxClient) GetHttpUrl(path APIPath) string {
	return wx.httpUrl + string(path)
}

// 获取发送主体消息
func GetRequestMsg(msg *MSG) (body string) {
	_b := &requetBody{
		Para: *msg,
	}
	b, _ := json.Marshal(_b)
	return string(b)
}

// ws通讯通用方法
func (wx *WxClient) Send(msg *MSG) (i int, err error) {
	m, err := json.Marshal(msg)
	if err != nil {
		log.Printf("msg:%v, %v, err: %v", msg, m, err)
		return
	}
	i, err = wx.ws.Write(m)
	return
}

func (wx *WxClient) parse(host string) {
	u := url.URL{
		Scheme: "ws",
		Host:   host,
	}
	wx.wsUrl = u.String()
	uu := (url.URL{
		Scheme: "http",
		Host:   host,
	})
	wx.httpUrl = uu.String()
}

// 动态读取大小
func (wx *WxClient) readMessage() (bytes []byte) {
again:
	fr, err := wx.ws.NewFrameReader()
	if err != nil {

		return
	}
	frame, err := wx.ws.HandleFrame(fr)
	if err != nil {
		return
	}
	if frame == nil {
		goto again
	}

	bytes, err = ioutil.ReadAll(frame)
	if err != nil {
		log.Printf("read frame data err %v", err)
	}
	return bytes
}
