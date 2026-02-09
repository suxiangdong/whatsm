package whats

import (
	"context"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) LoginCheck(ctx context.Context, req *v1.LoginCheckReq) (res *v1.LoginCheckRes, err error) {
	return &v1.LoginCheckRes{IsLogin: service.Whats().IsWhatsAccountLogin(ctx, req.Phone)}, nil
}
