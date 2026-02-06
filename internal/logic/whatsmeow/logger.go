package whatsmeow

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go.mau.fi/whatsmeow/util/log"
)

type logger struct {
	ctx context.Context
}

func (s *logger) Warnf(msg string, args ...interface{}) {
	g.Log().Warningf(s.ctx, msg, args...)
}

func (s *logger) Errorf(msg string, args ...interface{}) {
	g.Log().Errorf(s.ctx, msg, args...)
}

func (s *logger) Infof(msg string, args ...interface{}) {
	g.Log().Infof(s.ctx, msg, args...)
}

func (s *logger) Debugf(msg string, args ...interface{}) {
	g.Log().Debugf(s.ctx, msg, args...)
}

func (s *logger) Sub(_ string) waLog.Logger {
	return s
}
