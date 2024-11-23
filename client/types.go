package client

// 消息事件
type RespEvent struct {
	Code    int    `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

type Event struct {
	ID        string `json:"id"`
	ID1       string `json:"id1"` // 发送人
	ID2       string `json:"id2"` //
	ID3       string `json:"id3"`
	Type      int32  `json:"type"`
	Content   string `json:"content"`
	Receiver  string `json:"receiver"`
	Sender    string `json:"sender"`
	Srvid     int8   `json:"srvid"`
	Status    string `json:"status"`
	Time      string `json:"time"`
	Wxid      string `json:"wxid"` // 发送人 | 群ID
	ImagePath string `json:"image_path"`
	ThumbPath string `json:"thumb_path"`
	Title     string `json:"title"`
	Msg       string `json:"msg"`
}

// 图片消息事件 3
type EventPicConent struct {
	ID      string `json:"id"`
	Type    int32  `json:"type"`
	Content struct {
		Content string `json:"content"`
		Detail  string `json:"detail"` // 图片详情
		ID1     string `json:"id1"`    // 发送人 | 群ID
		ID2     string `json:"id2"`    //
		Thumb   string `json:"thumb"`  // 缩略图
	} `json:"content"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Srvid    int8   `json:"srvid"`
	Status   string `json:"status"`
	Time     string `json:"time"`
}

// 引用消息事件 49
type EventReferenceConent struct {
	ID      string `json:"id"`
	Type    int32  `json:"type"`
	Content struct {
		Content string `json:"content"`
		ID1     string `json:"id1"` // 发送人
		ID2     string `json:"id2"` //
	} `json:"content"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Srvid    int8   `json:"srvid"`
	Status   string `json:"status"`
	Time     string `json:"time"`
}

// user_list 5000
type EventUserList struct {
	ID      string `json:"id"`
	Type    int32  `json:"type"`
	Content []struct {
		Headimg string `json:"headimg"`
		Name    string `json:"name"`
		Node    int32  `json:"node"`
		Remarks string `json:"remarks"`
		Wxcode  string `json:"wxcode"` // 微信号
		Wxid    string `json:"wxid"`   // 微信id
	} `json:"content"`
}

// chatroom member 5010
type EventChatrooMmember struct {
	ID      string `json:"id"`
	Type    int32  `json:"type"`
	Content []struct {
		Address string   `json:"address"`
		Member  []string `json:"member"`
		RoomID  string   `json:"room_id"`
	} `json:"content"`
}

// chatroom member nickname 5020
type EventChatrooMmemberNick struct {
	ID       string `json:"id"`
	Type     int32  `json:"type"`
	Content  Nick   `json:"content"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Srvid    int8   `json:"srvid"`
	Status   string `json:"status"`
	Time     string `json:"time"`
}

type Nick struct {
	Nick   string `json:"nick"`
	Wxid   string `json:"wxid"`
	RoomID string `json:"roomid"`
}

// 拍一拍 || 添加好友
type EventPaiyipai struct {
	ID      string `json:"id"`
	Type    int32  `json:"type"`
	Content struct {
		Content string `json:"content"`
		ID1     string `json:"id1"` // 发送人 | 群ID
	} `json:"content"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	Srvid    int8   `json:"srvid"`
	Status   string `json:"status"`
	Time     string `json:"time"`
}

// 个人资料
type PersonalInfo struct {
	Code    string `json:"wxcode"`
	Id      string `json:"wxid"`
	Name    string `json:"nick"`
	HeadImg string `json:"head_img"`
}

// 消息主体
type MSG struct {
	Wxid    string `json:"wxid"`
	Content string `json:"content"`
}

// 内部使用解析content
type content struct {
	Content string `json:"content"`
}
