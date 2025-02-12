package user

import (
	"context"
	"serviceauth/internal/client/db"
	"serviceauth/internal/model"
	"serviceauth/internal/repository"
	"serviceauth/internal/repository/user/convertor"
	reposModel "serviceauth/internal/repository/user/model"
)

type repos struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repos{db: db}
}

func (r *repos) Get(ctx context.Context, id int) (*model.User, error) {
	if err := r.db.DB().Ping(ctx); err != nil {
		return &model.User{}, err
	}
	query := db.Query{
		Name:        "repository_get",
		QueryString: "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id=$1",
	}

	row, err := r.db.DB().QueryContext(ctx, query, id)
	if err != nil {
		return &model.User{}, err
	}
	var user reposModel.User
	for row.Next() {

		row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	}
	return convertor.ReposUserIntoServiceFromRepos(user), nil
}

func (r *repos) Create(ctx context.Context, user *model.UserInfo) (int, error) {
	if err := r.db.DB().Ping(ctx); err != nil {
		return 0, err
	}
	query := db.Query{
		Name:        "repository_create",
		QueryString: "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
	}

	row, err := r.db.DB().QueryContext(ctx, query, user.Name, user.Email, user.Password)
	var idU int
	for row.Next() {
		row.Scan(&idU)
	}
	if err != nil {
		return 0, err
	}
	return idU, nil
}

func (r *repos) Delete(ctx context.Context, id int) error {
	if err := r.db.DB().Ping(ctx); err != nil {
		return err
	}

	query := db.Query{
		Name:        "repository_delete",
		QueryString: "DELETE FROM users WHERE id=$1",
	}
	_, err := r.db.DB().ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repos) Update(ctx context.Context, user *model.UserUpdate) error {
	if err := r.db.DB().Ping(ctx); err != nil {
		return err
	}
	query := db.Query{
		Name:        "repository_update",
		QueryString: "UPDATE users SET name = $1, email = $2, updated_at = now() WHERE id=$3",
	}
	_, err := r.db.DB().ExecContext(ctx, query, user.Name, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}
