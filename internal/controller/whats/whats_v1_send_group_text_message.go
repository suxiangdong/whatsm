package whats

import (
	"context"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) SendGroupTextMessage(ctx context.Context, req *v1.SendGroupTextMessageReq) (res *v1.SendGroupTextMessageRes, err error) {
	if err := service.Whats().SendGroupTextMessage(ctx, req.From, req.To, req.Text); err != nil {
		return nil, err
	}
	return &v1.SendGroupTextMessageRes{}, nil
}
