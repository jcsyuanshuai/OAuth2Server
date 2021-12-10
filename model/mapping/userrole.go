package mapping

import "time"

type MapUserGroup struct {
	ID        int
	UserId    int
	RoleId    int
	Creator   string
	Updater   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
