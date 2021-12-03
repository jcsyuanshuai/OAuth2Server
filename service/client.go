package service

import (
	"github.com/go/oauth2-server/model/client"
)

func (cs *ClientStore) GetById() (client.Client, error) {
	return client.Client{}, nil
}
