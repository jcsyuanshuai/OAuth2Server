package role

import "time"

type Role struct {
	Id          int
	Name        string
	Creator     string
	Updater     string
	Scopes      string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
