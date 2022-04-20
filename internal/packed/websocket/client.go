package websocket

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/gorilla/websocket"
	"runtime/debug"
)

const (
	// 用户连接超时时间
	heartbeatExpirationTime = 6 * 60
)

// 用户登录
type login struct {
	UserId uint64
	Client *Client
}

// GetKey 读取客户端数据
func (l *login) GetKey() (key string) {
	key = GetUserKey(l.UserId)
	return
}

// Client 客户端连接
type Client struct {
	Addr          string          // 客户端地址
	ID            string          // 连接唯一标识
	Socket        *websocket.Conn // 用户连接
	Send          chan *WResponse // 待发送的数据
	SendClose     bool            // 发送是否关闭
	UserId        uint64          // 用户ID，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
	isApp         bool            // 是否是app
	tags          garray.StrArray // 标签
}

// NewClient 初始化
func NewClient(addr string, socket *websocket.Conn, firstTime uint64) (client *Client) {
	client = &Client{
		Addr:          addr,
		ID:            guid.S(),
		Socket:        socket,
		Send:          make(chan *WResponse, 100),
		SendClose:     false,
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
	}
	return
}

// 读取客户端数据
func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		c.close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			return
		}
		// 处理程序
		fmt.Println(message)
		ProcessData(c, message)
	}
}

// 向客户端写数据
func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()
	defer func() {
		clientManager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				// 发送数据错误 关闭连接
				return
			}
			_ = c.Socket.WriteJSON(message)
		}
	}
}

// SendMsg 发送数据
func (c *Client) SendMsg(msg *WResponse) {
	if c == nil || c.SendClose {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()
	c.Send <- msg
}

// Heartbeat 心跳更新
func (c *Client) Heartbeat(currentTime uint64) {
	c.HeartbeatTime = currentTime
	return
}

// IsHeartbeatTimeout 心跳是否超时
func (c *Client) IsHeartbeatTimeout(currentTime uint64) (timeout bool) {
	if c.HeartbeatTime+heartbeatExpirationTime <= currentTime {
		timeout = true
	}
	return
}

// 关闭客户端
func (c *Client) close() {
	if c.SendClose {
		return
	}
	c.SendClose = true
	close(c.Send)
}
