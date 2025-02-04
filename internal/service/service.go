package service

import (
	"context"
	"serviceauth/internal/model"
)

type UserService interface {
	Get(context.Context, int) (*model.User, error)
	Create(context.Context, *model.UserInfo) (int, error)
	Delete(context.Context, int) error
	Update(context.Context, *model.UserUpdate) error
}
