package token

import (
	"github.com/go/oauth2-server/model/client"
	"github.com/go/oauth2-server/model/user"
	"time"
)

type Token struct {
	User      user.User
	Client    client.Client
	Scopes    string
	Type      string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
