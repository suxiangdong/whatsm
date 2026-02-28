package hook

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"whatsm/internal/model"
	"whatsm/internal/service"
)

type sHook struct {
	c *gclient.Client
}

func init() {
	service.RegisterHook(New())
}

func New() service.IHook {
	return &sHook{
		c: gclient.New().ContentJson(),
	}
}

func (h *sHook) Trigger(ctx context.Context, data *model.HookData) error {
	gv, err := g.Cfg().Get(ctx, "callback.urls")
	if err != nil {
		return gerror.Wrap(err, "get callback.urls failed")
	}
	urls := gv.Strings()
	for _, url := range urls {
		if _, err := h.c.Post(ctx, url, data); err != nil {
			return gerror.Wrapf(err, "call back url: %s failed", url)
		}
	}
	//h.c.Post(ctx)
	return nil
}
