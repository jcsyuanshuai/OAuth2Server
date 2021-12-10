package token

import (
	"context"
	"github.com/go/oauth2-server/constant"
	pb "github.com/go/oauth2-server/proto/token"
)

func (s *Service) AuthorizationCode(ctx context.Context, req *pb.AuthorizationCodeReq) (*pb.TokenItem, error) {
	s.WithContext(ctx)

	return genToken(
		withGrantType(constant.GrantTypePasswordCredentials),
		withEmail(req.Email),
		withClientId(req.ClientId),
	)
}

func (s *Service) ClientCredential(ctx context.Context, req *pb.ClientCredentialReq) (*pb.TokenItem, error) {
	err := clientLogin(ctx, req.ClientId, req.ClientSecret)
	if err != nil {
		return nil, err
	}
	return genToken(
		withGrantType(constant.GrantTypeClientCredential),
		withClientId(req.ClientId),
	)
}

func (s *Service) PasswordCredential(ctx context.Context, req *pb.PasswordCredentialReq) (*pb.TokenItem, error) {
	err := userLogin(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return genToken(
		withGrantType(constant.GrantTypePasswordCredentials),
		withEmail(req.Email),
		withClientId(req.ClientId),
	)
}

func (s *Service) Implicit(ctx context.Context, req *pb.AuthorizationCodeReq) (*pb.TokenItem, error) {
	err := userLogin(ctx, req.Email, req.Password)
	if err != nil {
		return nil, err
	}
	return genToken(
		withGrantType(constant.GrantTypeImplicit),
		withEmail(req.Email),
		withClientId(req.ClientId),
	)
}
