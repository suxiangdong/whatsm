package v1

import "github.com/gogf/gf/v2/frame/g"

type SendTextMessageReq struct {
	g.Meta `path:"/whats/message/text/send" tags:"whats" method:"post" sm:"发送消息"`
	From   string `json:"from" dc:"手机号，带区号，不带加号"`
	To     string `json:"to" dc:"手机号，带区号，不带加号"`
	Text   string `json:"text" dc:"文字文本"`
}

type SendTextMessageRes struct{}

type SendGroupTextMessageReq struct {
	g.Meta `path:"/whats/message/group/text/send" tags:"whats" method:"post" sm:"发送消息到群组"`
	From   string `json:"from" dc:"手机号，带区号，不带加号"`
	To     string `json:"to" dc:"群组号，示例：120363123456789012"`
	Text   string `json:"text" dc:"文字文本"`
}

type SendGroupTextMessageRes struct{}
