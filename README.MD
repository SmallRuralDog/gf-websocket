# GoFrame websocket 的实现

## 实现功能

- 全局广播
- 单个房间推送
- 单个用户推送
- 单个客户端推送

## 客户端发送数据格式

```json
{
  "e": "event_name",
  "d": {
    "key": "value"
  }
}
```
