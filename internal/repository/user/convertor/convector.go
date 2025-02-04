package convertor

import (
	"serviceauth/internal/model"
	modelRepo "serviceauth/internal/repository/user/model"
)

func ReposUserIntoServiceFromRepos(user modelRepo.User) *model.User {
	return &model.User{
		Id: user.Id,
		UserInfo: model.UserInfo{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
		Created_at: user.Created_at,
		Updated_at: user.Updated_at,
	}
}
