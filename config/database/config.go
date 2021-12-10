package database

import (
	"context"
	"gorm.io/gorm"
	"sync"
)

const (
	MysqlEngineBase = ""
	RedisEngineBase = ""
)

var (
	databaseMap sync.Map
)

type Database struct {
	*gorm.DB
}

func (db *Database) SetupEngine(ctx context.Context, engine string) bool {
	if val, ok := databaseMap.Load(engine); ok {
		if item, ok := val.(*gorm.DB); ok {
			db.DB = item.WithContext(ctx)
			return true
		}
	}
	return false
}
