package model

import (
	"database/sql"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_user_v1"

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
	Role     auth_user_v1.Enum
}
