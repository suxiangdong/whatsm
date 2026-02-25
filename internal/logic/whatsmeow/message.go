package whatsmeow

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go.mau.fi/whatsmeow"
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

func (s *sWhats) SendGroupTextMessage(ctx context.Context, from, to, msg string) error {
	if _, ok := s.sessions[from]; !ok {
		return gerror.New("sender not found")
	}
	groupJID := types.NewJID(to, types.GroupServer)

	message := &waE2E.Message{
		Conversation: proto.String(msg),
	}
	resp, err := s.sessions[from].cli.SendMessage(ctx, groupJID, message)
	if err != nil {
		return gerror.Wrapf(err, "send group msg failed")
	}
	g.Log(consts.LogicLog).Debugf(ctx, "group message: send success, messageId: %s, timestamp: %v, from: %s, to: %s", resp.ID, resp.Timestamp, from, to)
	return nil
}

func buildMediaMessage(typ int, caption string, response *whatsmeow.UploadResponse) (*waE2E.Message, error) {
	switch typ {
	case consts.UploadFileImage:
		return &waE2E.Message{
			ImageMessage: &waE2E.ImageMessage{
				Caption:       proto.String(caption),
				URL:           proto.String(response.URL),
				DirectPath:    proto.String(response.DirectPath),
				MediaKey:      response.MediaKey,
				Mimetype:      proto.String("image/jpeg"),
				FileEncSHA256: response.FileEncSHA256,
				FileSHA256:    response.FileSHA256,
				FileLength:    proto.Uint64(response.FileLength),
			},
		}, nil
	default:
		return nil, gerror.New("media type not support, only image(1), video(2), audio(3)")
	}

}

func (s *sWhats) SendMediaMessage(ctx context.Context, from, to, caption string, typ int, rsp *whatsmeow.UploadResponse) error {
	targetJID := types.NewJID(to, types.DefaultUserServer)
	message, err := buildMediaMessage(typ, caption, rsp)
	if err != nil {
		return gerror.Wrap(err, "build media message failed")
	}
	if _, ok := s.sessions[from]; !ok {
		return gerror.New("sender not found")
	}
	resp, err := s.sessions[from].cli.SendMessage(ctx, targetJID, message)
	if err != nil {
		return gerror.Wrapf(err, "send msg failed")
	}
	g.Log(consts.LogicLog).Debugf(ctx, "mediaMessage: send success, messageId: %s, timestamp: %v, from: %s, to: %s", resp.ID, resp.Timestamp, from, to)
	return nil
}
