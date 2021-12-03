package client

import "time"

type Client struct {
	ID          string
	Secret      string
	RedirectUri string
	UserId      string
	AppName     string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

