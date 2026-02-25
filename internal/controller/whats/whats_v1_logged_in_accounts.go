package whats

import (
	"context"
	"whatsm/internal/service"

	"whatsm/api/whats/v1"
)

func (c *ControllerV1) LoggedInAccounts(ctx context.Context, req *v1.LoggedInAccountsReq) (res *v1.LoggedInAccountsRes, err error) {
	phones := service.Whats().LoggedInAccounts()
	as := make([]*v1.Account, 0)
	for _, phone := range phones {
		as = append(as, &v1.Account{Phone: phone})
	}
	return &v1.LoggedInAccountsRes{List: as}, nil
}
