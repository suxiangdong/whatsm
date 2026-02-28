package model

type HookData struct {
	Event   int    `json:"event" v:"required" dc:"事件，1 Login, 2 Logout"`
	Phone   string `json:"phone" dc:"账号"`
	Message string `json:"message" dc:"事件信息"`
}
