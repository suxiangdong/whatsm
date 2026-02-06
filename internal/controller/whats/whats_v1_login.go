package whats

import (
	"context"
	"whatsm/internal/model"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	out, err := service.Whats().LoginPair(ctx, &model.LoginPairInput{Proxy: req.Proxy, Phone: req.Phone})
	if err != nil {
		return nil, err
	}
	return &v1.LoginRes{Code: out.Code, QrCode: out.QrCode}, nil
}
