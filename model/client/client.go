package client

import "time"

type Client struct {
	Id          string
	Secret      string
	RedirectUri string
	UserId      string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
