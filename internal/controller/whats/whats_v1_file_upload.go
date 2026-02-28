package whats

import (
	"context"
	"encoding/hex"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) FileUpload(ctx context.Context, req *v1.FileUploadReq) (res *v1.FileUploadRes, err error) {
	rsp, err := service.Whats().Upload(ctx, req.From, req.Type, req.File)
	if err != nil {
		return nil, err
	}
	return &v1.FileUploadRes{
		MimeType:      rsp.MimeType,
		URL:           rsp.Rsp.URL,
		DirectPath:    rsp.Rsp.DirectPath,
		MediaKey:      hex.EncodeToString(rsp.Rsp.MediaKey),
		FileSHA256:    hex.EncodeToString(rsp.Rsp.FileSHA256),
		FileEncSHA256: hex.EncodeToString(rsp.Rsp.FileEncSHA256),
		FileLength:    rsp.Rsp.FileLength,
	}, nil
}
