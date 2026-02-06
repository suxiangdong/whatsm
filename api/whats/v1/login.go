package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta `path:"/whats/login" tags:"whats" method:"post" sm:"登录"`
	Phone  string `json:"phone" v:"required" dc:"手机号，带区号，不带+号"`
	Proxy  string `json:"proxy" dc:"代理地址，示例：socks5://root:aasd1123@127.0.0.1:10808"`
}

type LoginRes struct {
	Code   string `json:"code" dc:"pair code"`
	QrCode string `json:"qrCode" dc:"二维码"`
}
