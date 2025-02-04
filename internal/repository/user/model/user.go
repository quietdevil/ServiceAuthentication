package model

import "time"

type User struct {
	Id int
	UserInfo
	Created_at time.Time
	Updated_at time.Time
}

type UserInfo struct {
	Name     string
	Email    string
	Password string
}
