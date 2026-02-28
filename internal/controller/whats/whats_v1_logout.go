package whats

import (
	"context"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return &v1.LogoutRes{}, service.Whats().Logout(ctx, req.Phone)
}
