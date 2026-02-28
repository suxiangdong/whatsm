package whats

import (
	"context"
	"fmt"
	"whatsm/api/whats/v1"
)

func (c *ControllerV1) Hook(ctx context.Context, req *v1.HookReq) (res *v1.HookRes, err error) {
	fmt.Println(req.Event, req.Message, req.Phone)
	return &v1.HookRes{}, nil
}
