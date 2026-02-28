package v1

import "github.com/gogf/gf/v2/frame/g"

type SendTextMessageReq struct {
	g.Meta `path:"/whats/message/text/send" tags:"whats" method:"post" sm:"发送消息"`
	From   string `json:"from" v:"required" dc:"手机号，带区号，不带加号"`
	To     string `json:"to" v:"required" dc:"手机号，带区号，不带加号"`
	Text   string `json:"text" v:"required" dc:"文字文本"`
}

type SendTextMessageRes struct{}

type SendGroupTextMessageReq struct {
	g.Meta `path:"/whats/message/group/text/send" tags:"whats" method:"post" sm:"发送消息到群组"`
	From   string `json:"from" v:"required" dc:"手机号，带区号，不带加号"`
	To     string `json:"to" v:"required" dc:"群组号，示例：120363123456789012"`
	Text   string `json:"text" v:"required" dc:"文字文本"`
}

type SendGroupTextMessageRes struct{}

type SendMediaMessageReq struct {
	g.Meta        `path:"/whats/message/media/send" tags:"whats" method:"post" sm:"发送媒体消息"`
	MimeType      string `json:"mimeType" v:"required" dc:"mime_type"`
	Type          int    `json:"type" v:"required" dc:"媒体类型，1图片，2视频，3音频，当前仅支持图片"`
	From          string `json:"from" v:"required" dc:"手机号，带区号，不带加号"`
	To            string `json:"to" v:"required" dc:"群组号，示例：120363123456789012"`
	Caption       string `json:"caption" v:"required" dc:"媒体描述"`
	URL           string `json:"url" v:"required" dc:"上传文件接口返回的数据"`
	DirectPath    string `json:"directPath" v:"required" dc:"上传文件接口返回的数据"`
	MediaKey      string `json:"mediaKey" v:"required" dc:"上传文件接口返回的数据"`
	FileEncSHA256 string `json:"fileEncSHA256" v:"required" dc:"上传文件接口返回的数据"`
	FileSHA256    string `json:"fileSHA256" v:"required" dc:"上传文件接口返回的数据"`
	FileLength    uint64 `json:"fileLength" v:"required" dc:"上传文件接口返回的数据"`
}

type SendMediaMessageRes struct{}
