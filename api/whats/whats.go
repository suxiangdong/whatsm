// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package whats

import (
	"context"

	"whatsm/api/whats/v1"
)

type IWhatsV1 interface {
	FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	LoginCheck(ctx context.Context, req *v1.LoginCheckReq) (res *v1.LoginCheckRes, err error)
	LoggedInAccounts(ctx context.Context, req *v1.LoggedInAccountsReq) (res *v1.LoggedInAccountsRes, err error)
	SendTextMessage(ctx context.Context, req *v1.SendTextMessageReq) (res *v1.SendTextMessageRes, err error)
	SendGroupTextMessage(ctx context.Context, req *v1.SendGroupTextMessageReq) (res *v1.SendGroupTextMessageRes, err error)
	SendMediaMessage(ctx context.Context, req *v1.SendMediaMessageReq) (res *v1.SendMediaMessageRes, err error)
}
