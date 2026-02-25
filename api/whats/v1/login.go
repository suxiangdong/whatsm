package v1

import "github.com/gogf/gf/v2/frame/g"

type LoginReq struct {
	g.Meta `path:"/whats/login" tags:"whats" method:"post" sm:"登录" dc:"用户登录，最多同时登录N（默认200，可以在应用的配置文件设置）个账号。本接口有两个特殊code码，1000代表被whats官方限制（overlimit），1001代表登录的账号数量达到上限"`
	Phone  string `json:"phone" v:"required" dc:"手机号，带区号，不带+号"`
	Proxy  string `json:"proxy" dc:"代理地址，示例：socks5://root:aasd1123@127.0.0.1:10808"`
}

type LoginRes struct {
	Code   string `json:"code" dc:"pair code"`
	QrCode string `json:"qrCode" dc:"二维码"`
}

type LoginCheckReq struct {
	g.Meta `path:"/whats/login_check" tags:"whats" method:"get" sm:"检查登录" dc:"根据手机号检查账号是否登录"`
	Phone  string `json:"phone" v:"required" dc:"手机号，带区号，不带+号"`
}

type LoginCheckRes struct {
	IsLogin bool `json:"isLogin" dc:"是否登录"`
}

type LoggedInAccountsReq struct {
	g.Meta `path:"/whats/logged_accounts" tags:"whats" method:"get" sm:"已登录账号" dc:"获取所有已登录账号"`
}

type Account struct {
	Phone string `json:"phone"`
}

type LoggedInAccountsRes struct {
	List []*Account `json:"list" dc:"登录账号列表"`
}
