package websocket

import "github.com/gogf/gf/v2/frame/g"

// 当前输入对象
type request struct {
	Event string `json:"e"` //事件名称
	Data  g.Map  `json:"d"` //数据
}

// WResponse 输出对象
type WResponse struct {
	Event string      `json:"e"` //事件名称
	Data  interface{} `json:"d"` //数据
}

type TagWResponse struct {
	Tag       string
	WResponse *WResponse
}

type UserWResponse struct {
	UserID    uint64
	WResponse *WResponse
}

type ClientWResponse struct {
	ID        string
	WResponse *WResponse
}
