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
	IHook interface {
		Trigger(ctx context.Context, data *model.HookData) error
	}
)

var (
	localHook IHook
)

func Hook() IHook {
	if localHook == nil {
		panic("implement not found for interface IHook, forgot register?")
	}
	return localHook
}

func RegisterHook(i IHook) {
	localHook = i
}
