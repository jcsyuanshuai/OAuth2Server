package token

import (
	"context"
	pb "github.com/go/oauth2-server/proto/token"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) Authorize(ctx context.Context, req *pb.AuthorizeReq) (_ *emptypb.Empty, err error) {
	s.WithContext(ctx)
	info := &pb.TokenInfo{}
	err = Decode(req.Token, info)
	if err != nil {
		return nil, err
	}
	err = userExists(ctx, info.Email)
	if err != nil {
		return nil, err
	}
	err = clientExists(ctx, info.ClientId)
	if err != nil {
		return nil, err
	}
	return
}
