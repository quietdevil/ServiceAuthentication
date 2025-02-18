package authentication

import (
	"context"
	"fmt"
	"github.com/quietdevil/Platform_common/pkg/db"
	"serviceauth/internal/model"
	"serviceauth/internal/repository/user/convertor"
	reposModel "serviceauth/internal/repository/user/model"
)

func (r *AuthRepos) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if err := r.DB.DB().Ping(ctx); err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "repository_get_by_username",
		QueryStr: "SELECT id, name, email, password, created_at, updated_at FROM users WHERE name=$1",
	}

	row, err := r.DB.DB().ContextQuery(ctx, query, username)
	if err != nil {
		return nil, err
	}
	var user reposModel.User
	for row.Next() {

		row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	}
	fmt.Println(user.Password)
	return convertor.ReposUserIntoServiceFromRepos(user), nil

}
