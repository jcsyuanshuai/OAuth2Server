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
	Department  string
	Position    string
	Level       string
	LastLoginAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
