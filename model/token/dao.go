package token

import (
	"context"
	"github.com/go/oauth2-server/config/database"
)

type Dao struct {
	ctx context.Context
	db  *database.Database
}

func New(ctx context.Context) *Dao {
	db := &database.Database{}
	db.SetupEngine(ctx, database.RedisEngineBase)
	return &Dao{
		ctx: ctx,
		db:  db,
	}
}
