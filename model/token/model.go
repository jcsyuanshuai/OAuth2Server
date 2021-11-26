package token

import "time"

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
	Scopes    []string
	ExpiresIn time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
