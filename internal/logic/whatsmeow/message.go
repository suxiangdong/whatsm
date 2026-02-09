package whatsmeow

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
	"whatsm/internal/consts"
)

func (s *sWhats) SendTextMessage(ctx context.Context, from, to, msg string) error {
	targetJID := types.NewJID(to, types.DefaultUserServer)

	message := &waE2E.Message{
		Conversation: proto.String(msg),
	}
	if _, ok := s.sessions[from]; !ok {
		return gerror.New("sender not found")
	}

	resp, err := s.sessions[from].cli.SendMessage(ctx, targetJID, message)
	if err != nil {
		return gerror.Wrapf(err, "send msg failed")
	}
	g.Log(consts.LogicLog).Debugf(ctx, "message: send success, messageId: %s, timestamp: %v, from: %s, to: %s", resp.ID, resp.Timestamp, from, to)
	return nil
}
