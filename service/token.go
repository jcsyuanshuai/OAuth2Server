package service

import "github.com/go/oauth2-server/model"

func (ts *TokenStore) Create() (model.Token, error) {
	return model.Token{}, nil
}

func (ts *TokenStore) Get(tokenType string, token string) (model.Token, error) {
	return model.Token{}, nil
}

func (ts *TokenStore) Remove(tokenType string, token string) error {
	return nil
}
