package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/whats/file" mime:"multipart/form-data" tags:"whats" method:"post" sm:"上传文件" dc:"上传文件"`
	File   *ghttp.UploadFile `json:"file" type:"file" dc:"选择上传文件"`
	From   string            `json:"from" v:"required" dc:"手机号，带区号，不带+号"`
	Type   int               `json:"type" v:"required" dc:"1 图片 2 视频 3音频"`
}

type FileUploadRes struct {
	URL           string `json:"url"`
	DirectPath    string `json:"directPath"`
	MediaKey      string `json:"mediaKey"`
	FileEncSHA256 string `json:"fileEncSHA256"`
	FileSHA256    string `json:"fileSHA256"`
	FileLength    uint64 `json:"fileLength"`
}
