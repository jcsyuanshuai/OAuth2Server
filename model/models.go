package model

import "time"

type User struct {
	ID          int
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Avatar      string
	Status      int32
	Position    string
	Level       string
	LastLoginAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Group struct {
	ID          int
	Name        string
	Creator     string
	Updater     string
	Scopes      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MapUserGroup struct {
	ID        int
	UserID    int
	GroupID   int
	Creator   string
	Updater   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Permission struct {
	ID   int
	Name string
}

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

const (
	TokenTypeAccess   = "Access"
	TokenTypeAuthCode = "Code"
	TokenTypeRefresh  = "Refresh"
)

type Model struct {
	ID        string
	UserID    string
	ClientID  string
	Type      string
	Token     string
	Scopes    string
	ExpiresIn time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
