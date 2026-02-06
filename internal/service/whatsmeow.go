// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"whatsm/internal/model"
)

type (
	IWhats interface {
		SendTextMessage(ctx context.Context, from string, to string, msg string) error
		// Init connect to db
		Init(ctx context.Context) error
		// create new device&session
		LoginPair(ctx context.Context, in *model.LoginPairInput) (*model.LoginPairOutput, error)
	}
)

var (
	localWhats IWhats
)

func Whats() IWhats {
	if localWhats == nil {
		panic("implement not found for interface IWhats, forgot register?")
	}
	return localWhats
}

func RegisterWhats(i IWhats) {
	localWhats = i
}
