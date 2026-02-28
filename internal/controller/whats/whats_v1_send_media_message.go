package whats

import (
	"context"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"go.mau.fi/whatsmeow"
	"whatsm/internal/model"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) SendMediaMessage(ctx context.Context, req *v1.SendMediaMessageReq) (res *v1.SendMediaMessageRes, err error) {

	mediaKey, err := hex.DecodeString(req.MediaKey)
	if err != nil {
		return nil, gerror.Wrap(err, "hex decode mediaKey failed")
	}
	fileSha256, err := hex.DecodeString(req.FileSHA256)
	if err != nil {
		return nil, gerror.Wrap(err, "hex decode fileSha256 failed")
	}
	fileEncSha256, err := hex.DecodeString(req.FileEncSHA256)
	if err != nil {
		return nil, gerror.Wrap(err, "hex decode fileEncSha256 failed")
	}
	if err := service.Whats().SendMediaMessage(ctx, &model.SendMediaMessageInput{
		From:     req.From,
		To:       req.To,
		Type:     req.Type,
		Caption:  req.Caption,
		MimeType: req.MimeType,
		Rsp: &whatsmeow.UploadResponse{
			URL:           req.URL,
			DirectPath:    req.DirectPath,
			MediaKey:      mediaKey,
			FileSHA256:    fileSha256,
			FileEncSHA256: fileEncSha256,
			FileLength:    req.FileLength,
		},
	}); err != nil {
		return nil, err
	}
	return &v1.SendMediaMessageRes{}, nil
}
