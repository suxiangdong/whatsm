package whats

import (
	"context"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) SendTextMessage(ctx context.Context, req *v1.SendTextMessageReq) (res *v1.SendTextMessageRes, err error) {
	if err := service.Whats().SendTextMessage(ctx, req.From, req.To, req.Text); err != nil {
		return nil, err
	}
	return &v1.SendTextMessageRes{}, nil
}
