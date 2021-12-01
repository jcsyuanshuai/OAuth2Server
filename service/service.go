package service

import (
	"github.com/go/oauth2-server/model"
	"gorm.io/gorm"
)

type ClientService interface {
	GetById(id string) (model.Client, error)
}

type TokenService interface {
	Create() (model.Token, error)
	Remove(tokenType string, token string) error
	Get(tokenType string, token string) (model.Token, error)
}

type TokenStore struct {
	db *gorm.DB
}

type ClientStore struct {
	db *gorm.DB
}

type UserStore struct {

}
