package user

import (
	"context"
	"serviceauth/internal/model"
	"serviceauth/internal/repository/user/convertor"
	reposModel "serviceauth/internal/repository/user/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repos struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repos {
	return &repos{db: db}
}

func (r *repos) Get(ctx context.Context, id int) (*model.User, error) {
	if err := r.db.Ping(ctx); err != nil {
		return &model.User{}, err
	}
	row, err := r.db.Query(ctx, "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id=$1", id)

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
	if err := r.db.Ping(ctx); err != nil {
		return 0, err
	}
	_, err := r.db.Exec(ctx, "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (r *repos) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, "DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repos) Update(ctx context.Context, user *model.UserUpdate) error {
	_, err := r.db.Exec(ctx, "UPDATE users SET name = $1, email = $2, updated_at = now() WHERE id=$3", user.Name, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}
