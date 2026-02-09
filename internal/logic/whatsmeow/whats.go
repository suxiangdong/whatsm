package whatsmeow

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"sync"
	"whatsm/internal/consts"
	"whatsm/internal/model"
	"whatsm/internal/service"
)

type sWhats struct {
	c   *sqlstore.Container
	l   *logger
	ctx context.Context

	mu       sync.Mutex
	sessions map[string]*session
}

func init() {
	service.RegisterWhats(New())
}

func New() service.IWhats {
	return &sWhats{sessions: make(map[string]*session), mu: sync.Mutex{}}
}

// 检查账号是否登录
func (s *sWhats) IsWhatsAccountLogin(ctx context.Context, phone string) bool {
	sess, ok := s.sessions[phone]
	if !ok {
		return false
	}
	return sess.cli.IsLoggedIn()
}

// Init connect to db
func (s *sWhats) Init(ctx context.Context) error {
	s.ctx = ctx
	s.l = &logger{ctx: ctx}
	dialect := consts.DbDialectDefault
	address := consts.DbAddressDefault
	if dialectCfg, err := g.Cfg().Get(ctx, consts.DbDialectConfigKey); err == nil {
		dialect = dialectCfg.String()
	}
	if addressCfg, err := g.Cfg().Get(ctx, consts.DbAddressConfigKey); err == nil {
		address = addressCfg.String()
	}
	container, err := sqlstore.New(ctx, dialect, address, s.l)
	if err != nil {
		return gerror.Wrapf(err, "connect to db failed")
	}
	s.c = container
	return nil
}

// create new device&session
func (s *sWhats) LoginPair(ctx context.Context, in *model.LoginPairInput) (*model.LoginPairOutput, error) {
	limit := consts.MaxUserDefault
	if maxUser, err := g.Cfg().Get(ctx, consts.AutoMarkMessageKey); err == nil {
		if maxUser.Int() != 0 {
			limit = maxUser.Int()
		}
	}
	if len(s.sessions) >= limit {
		return nil, gerror.NewCode(gcode.New(1001, "login users limit", nil))
	}
	jid := types.NewADJID(in.Phone, 0, 1)
	st, err := s.c.GetDevice(ctx, jid)
	if err != nil {
		return nil, gerror.Wrapf(err, "device not found")
	}
	if st == nil {
		st = s.c.NewDevice()
	}
	client := whatsmeow.NewClient(st, s.l)
	if in.Proxy != "" {
		if err := client.SetProxyAddress(in.Proxy); err != nil {
			return nil, gerror.Wrapf(err, "set proxy address failed")
		}
	}
	sess := &session{cli: client, sw: s}
	autoMarkMessage := false
	if mark, err := g.Cfg().Get(ctx, consts.AutoMarkMessageKey); err == nil {
		autoMarkMessage = mark.Bool()
	}
	if autoMarkMessage {
		sess.hooks = []EventHook{HookMarkMessageAdRead}
	}
	client.AddEventHandler(sess.eventHandler)

	if client.Store.ID != nil {
		if err := client.Connect(); err != nil {
			return nil, gerror.Wrapf(err, "client connect to whats server failed")
		}
		return &model.LoginPairOutput{}, nil
	}
	client.Store.Platform = consts.PlatformDefault
	if pfCfg, err := g.Cfg().Get(ctx, consts.PlatformConfigKey); err == nil {
		client.Store.Platform = pfCfg.String()
	}
	client.Store.BusinessName = consts.BusinessNameDefault
	if bnCfg, err := g.Cfg().Get(ctx, consts.BusinessNameConfigKey); err == nil {
		client.Store.BusinessName = bnCfg.String()
	}
	client.Store.PushName = consts.PushNameDefault
	if pnCfg, err := g.Cfg().Get(ctx, consts.PushNameConfigKey); err == nil {
		client.Store.PushName = pnCfg.String()
	}
	qrChan, _ := client.GetQRChannel(context.Background())
	if err := client.Connect(); err != nil {
		return nil, gerror.Wrapf(err, "client dial to whats server failed")
	}
	// ensure websocket is ok
	qrCode := <-qrChan

	code, err := client.PairPhone(ctx, in.Phone, true, whatsmeow.PairClientChrome, consts.ClientDisplayNameDefault)
	if err != nil {
		if errors.Is(err, whatsmeow.ErrIQRateOverLimit) {
			return nil, gerror.NewCode(gcode.New(1000, err.Error(), nil))
		}
		return nil, gerror.Wrapf(err, "create pair code failed")
	}
	// 等待连接成功
	return &model.LoginPairOutput{Code: code, QrCode: qrCode.Code}, nil
}

//func (s *sWhats) login(ctx context.Context, st *store.Device, proxy ...string) error {
//
//}
