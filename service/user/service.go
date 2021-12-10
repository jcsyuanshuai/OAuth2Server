package user

import (
	"context"
	"github.com/go/oauth2-server/model/user"
	pb "github.com/go/oauth2-server/proto/user"
)

type Service struct {
	pb.UnimplementedUserServiceServer
	dao *user.Dao
}

func New() *Service {
	return &Service{}
}

var _ pb.UserServiceServer = &Service{}

func (s *Service) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	s.WithContext(ctx)
	return s.dao.Create(req)
}

func (s *Service) WithContext(ctx context.Context) {
	s.dao = user.New(ctx)
}
