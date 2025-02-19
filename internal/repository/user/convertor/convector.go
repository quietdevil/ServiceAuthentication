package convertor

import (
	"github.com/quietdevil/ServiceAuthentication/internal/model"
	modelRepo "github.com/quietdevil/ServiceAuthentication/internal/repository/user/model"
)

func ReposUserIntoServiceFromRepos(user modelRepo.User) *model.User {
	return &model.User{
		Id: user.Id,
		UserInfo: model.UserInfo{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Role:     user.Role.String(),
		},
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
