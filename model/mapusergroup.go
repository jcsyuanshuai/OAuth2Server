package model

import "time"

type MapUserGroup struct {
	ID        int
	UserID    int
	GroupID   int
	Creator   string
	Updater   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
