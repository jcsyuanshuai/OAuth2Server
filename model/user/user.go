package user

import "time"

type User struct {
	Id          int64
	No          string
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Avatar      string
	Status      int32
	Department  string
	Position    string
	Level       string
	LastLoginAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
