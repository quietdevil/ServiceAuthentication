package model

import (
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Id int
	UserInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Name     string
	Email    string
	Password string
	Role     string
}

type UserUpdate struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

type UserLogin struct {
	Username string
	Password string
}
