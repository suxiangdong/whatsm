package model

type LoginPairInput struct {
	Phone string
	Proxy string
}

type LoginPairOutput struct {
	QrCode string
	Code   string
}
