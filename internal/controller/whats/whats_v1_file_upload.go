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
		URL:           rsp.URL,
		DirectPath:    rsp.DirectPath,
		MediaKey:      hex.EncodeToString(rsp.MediaKey),
		FileSHA256:    hex.EncodeToString(rsp.FileSHA256),
		FileEncSHA256: hex.EncodeToString(rsp.FileEncSHA256),
		FileLength:    rsp.FileLength,
	}, nil
}
