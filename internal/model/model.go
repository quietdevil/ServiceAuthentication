package model

import "time"

type User struct {
	Id         int
	UserInfo   UserInfo
	Created_at time.Time
	Updated_at time.Time
}

type UserInfo struct {
	Name     string
	Email    string
	Password string
}

type UserUpdate struct {
	Id       int
	Name     string
	Email    string
	Password string
}
