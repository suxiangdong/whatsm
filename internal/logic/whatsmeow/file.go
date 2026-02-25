package whatsmeow

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"go.mau.fi/whatsmeow"
	"io"
	"whatsm/internal/consts"
)

func (s *sWhats) Upload(ctx context.Context, from string, typ int, uf *ghttp.UploadFile) (*whatsmeow.UploadResponse, error) {
	mediaType := whatsmeow.MediaType("")
	switch typ {
	case consts.UploadFileImage:
		mediaType = whatsmeow.MediaImage
	case consts.UploadFileVideo:
		mediaType = whatsmeow.MediaVideo
	case consts.UploadFileAudio:
		mediaType = whatsmeow.MediaAudio
	default:
		return nil, gerror.New("media type not support, only image(1), video(2), audio(3)")
	}
	if _, ok := s.sessions[from]; !ok {
		return nil, gerror.New("sender not found")
	}
	file, err := uf.Open()
	if err != nil {
		return nil, gerror.Wrap(err, `UploadFile.Open failed`)
	}
	defer file.Close()
	c, err := io.ReadAll(file)
	if err != nil {
		return nil, gerror.Wrap(err, "io.ReadAll failed")
	}
	resp, err := s.sessions[from].cli.Upload(ctx, c, mediaType)
	if err != nil {
		return nil, gerror.Wrap(err, "WhatsApp upload file failed")
	}
	return &resp, nil
}
