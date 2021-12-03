package service

import (
	"github.com/go/oauth2-server/model/auth"
	"github.com/go/oauth2-server/model/client"
	"gorm.io/gorm"
)

type ClientService interface {
	GetById(id string) (client.Client, error)
}

type TokenService interface {
	Create() (token.Token, error)
	Remove(tokenType string, token string) error
	Get(tokenType string, token string) (token.Token, error)
}

type TokenStore struct {
	db *gorm.DB
}

type ClientStore struct {
	db *gorm.DB
}

type UserStore struct {

}
