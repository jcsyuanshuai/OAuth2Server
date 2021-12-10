package token

import (
	"context"
	"github.com/go/oauth2-server/constant"
	"github.com/go/oauth2-server/model/token"
	pb "github.com/go/oauth2-server/proto/token"
	"time"
)

type Service struct {
	pb.UnimplementedTokenServiceServer
	dao *token.Dao
}

func New() *Service {
	return &Service{}
}

var _ pb.TokenServiceServer = &Service{}

var (
	DefaultExpiresIn int64 = 48 * 60 * 60
)

func (s *Service) WithContext(ctx context.Context) {
	s.dao = token.New(ctx)
}

func clientExists(ctx context.Context, key string) error {
	return nil
}

func userExists(ctx context.Context, email string) error {
	return nil
}

func clientLogin(ctx context.Context, key string, secret string) error {
	return nil
}

func userLogin(ctx context.Context, email string, password string) error {
	return nil
}

type opt func(token *pb.TokenInfo)

func genToken(opts ...opt) (ret *pb.TokenItem, err error) {
	info := &pb.TokenInfo{
		ExpiresIn: DefaultExpiresIn,
		CreatedAt: time.Now().Format(constant.TimeFormat),
	}
	for _, o := range opts {
		o(info)
	}
	ret = &pb.TokenItem{
		ExpiresIn: DefaultExpiresIn,
	}
	info.TokenType = ""
	ret.AccessToken, err = Encode(info)
	info.TokenType = ""
	ret.RefreshToken, err = Encode(info)

	return ret, nil
}

func Encode(data interface{}) (string, error) {
	return "", nil
}

func Decode(str string, dst interface{}) error {
	return nil
}

func withGrantType(grantType string) opt {
	return func(token *pb.TokenInfo) {
		token.GrantType = grantType
	}
}

func withEmail(grantType string) opt {
	return func(token *pb.TokenInfo) {
		token.GrantType = grantType
	}
}

func withClientId(grantType string) opt {
	return func(token *pb.TokenInfo) {
		token.GrantType = grantType
	}
}

func withTokenType(grantType string) opt {
	return func(token *pb.TokenInfo) {
		token.GrantType = grantType
	}
}

func withScopes(grantType string) opt {
	return func(token *pb.TokenInfo) {
		token.GrantType = grantType
	}
}
