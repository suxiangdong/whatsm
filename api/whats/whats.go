// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package whats

import (
	"context"

	"whatsm/api/whats/v1"
)

type IWhatsV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	SendTextMessage(ctx context.Context, req *v1.SendTextMessageReq) (res *v1.SendTextMessageRes, err error)
}
