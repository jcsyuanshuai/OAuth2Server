package user

import (
	"context"
	"github.com/go/oauth2-server/config/database"
	"github.com/go/oauth2-server/proto/user"
	"time"
)

type Dao struct {
	ctx context.Context
	db  *database.Database
}

func New(ctx context.Context) *Dao {
	db := &database.Database{}
	db.SetupEngine(ctx, database.MysqlEngineBase)
	return &Dao{
		ctx: ctx,
		db:  db,
	}
}

func (d *Dao) Create(req *user.User) (*user.User, error) {
	now := time.Now()
	tx := d.db.Create(&User{
		Id:          req.Id,
		No:          "",
		Name:        req.Name,
		Email:       req.Email,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Avatar:      req.Avatar,
		Status:      req.Status,
		Department:  req.Department,
		Position:    req.Position,
		Level:       req.Level,
		LastLoginAt: time.Time{},
		CreatedAt:   now,
		UpdatedAt:   now,
	})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return nil, nil
}
