package authentication

import (
	"context"
	"fmt"
	"github.com/quietdevil/Platform_common/pkg/db"
	"serviceauth/internal/model"
)

func (r *AuthRepos) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	if err := r.DB.DB().Ping(ctx); err != nil {
		return nil, err
	}

	query := db.Query{
		Name:     "repository_get_by_username",
		QueryStr: "SELECT * FROM users WHERE name=$1",
	}

	row, err := r.DB.DB().ContextQuery(ctx, query, username)
	if err != nil {
		return nil, err
	}
	var user model.User

	for row.Next() {

		row.Scan(&user.Id, &user.UserInfo.Name, &user.UserInfo.Email, &user.UserInfo.Password, &user.UserInfo.Role, &user.CreatedAt, &user.UpdatedAt)
	}
	fmt.Println(user)
	return &user, nil

}
