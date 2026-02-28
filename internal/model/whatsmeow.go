package model

import "go.mau.fi/whatsmeow"

type LoginPairInput struct {
	Phone string
	Proxy string
}

type LoginPairOutput struct {
	QrCode string
	Code   string
}

type SendMediaMessageInput struct {
	From     string
	To       string
	Caption  string
	MimeType string
	Type     int
	Rsp      *whatsmeow.UploadResponse
}

type UploadOutput struct {
	Rsp      *whatsmeow.UploadResponse
	MimeType string
}
