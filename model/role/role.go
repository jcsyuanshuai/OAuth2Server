package role

import "time"

type Role struct {
	Id          int
	Name        string
	Creator     string
	Updater     string
	Scopes      string //[1,2,3] permissions
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
