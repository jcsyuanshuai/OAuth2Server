package model

import "time"

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
