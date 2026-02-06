package whatsmeow

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	"whatsm/internal/consts"
)

type EventHook func(ctx context.Context, cli *whatsmeow.Client, msg *events.Message)

func HookMarkMessageAdRead(ctx context.Context, cli *whatsmeow.Client, msg *events.Message) {
	ids := []types.MessageID{msg.Info.ID}
	err := cli.MarkRead(ctx, ids, msg.Info.Timestamp, msg.Info.Chat, msg.Info.Sender)
	if err != nil {
		g.Log(consts.LogicLog).Debugf(ctx, "mark msg %s as read failed, err: %s", msg.Info.ID, err)
		return
	}
	g.Log(consts.LogicLog).Debugf(ctx, "msg %s marked as read", msg.Info.ID)
}
