package token

import (
	"context"
	"github.com/go/oauth2-server/constant"
	pb "github.com/go/oauth2-server/proto/token"
)

func (s *Service) Refresh(ctx context.Context, req *pb.RefreshReq) (*pb.TokenItem, error) {
	s.WithContext(ctx)
	info := &pb.TokenInfo{}
	err := Decode(req.Token, info)
	if err != nil {
		return nil, err
	}
	return genToken(
		withGrantType(constant.GrantTypePasswordCredentials),
		withEmail(info.Email),
		withClientId(info.ClientId),
	)
}
