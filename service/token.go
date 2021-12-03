package service

import (
	"github.com/go/oauth2-server/model/auth"
)

func (ts *TokenStore) Create() (token.Token, error) {
	return token.Token{}, nil
}

func (ts *TokenStore) Get(tokenType string, token string) (token.Token, error) {
	return token.Token{}, nil
}

func (ts *TokenStore) Remove(tokenType string, token string) error {
	return nil
}
