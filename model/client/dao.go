package client

import (
	"context"
	"gorm.io/gorm"
)

type Dao struct {
	ctx context.Context
	db  gorm.DB
}

func New() *Dao {
	return &Dao{}
}

func (dao *Dao) Create() *Client {
	//dao.db.Create()
	return nil
}
