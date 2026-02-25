package whatsmeow

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"go.mau.fi/whatsmeow/types"
)

func (s *sWhats) GetContactInfo(ctx context.Context, u, t string) error {
	if _, ok := s.sessions[u]; !ok {
		return gerror.New("sender not found")
	}
	targetJID := types.NewJID(t, types.DefaultUserServer)

	// 获取联系人信息
	info, err := s.sessions[u].cli.Store.Contacts.GetContact(ctx, targetJID)
	if err != nil {
		fmt.Printf("获取联系人信息失败: %v\n", err)
		return gerror.Wrapf(err, "get contact failed")
	}
	fmt.Printf("联系人信息:\n")
	fmt.Printf("  名称: %s\n", info.FullName)
	fmt.Printf("  昵称: %s\n", info.PushName)
	fmt.Printf("  商业名称: %s\n", info.BusinessName)
	return nil
}
