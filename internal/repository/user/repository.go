package user

import (
	"context"
	"serviceauth/internal/model"
	"serviceauth/internal/repository"

	db "github.com/quietdevil/Platform_common/pkg/db"
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
		Name:     "repository_get",
		QueryStr: "SELECT id, name, email, password,roles,created_at, updated_at FROM users WHERE id=$1",
	}

	//switch a.(type) {
	//case string:
	//	query.QueryStr = "SELECT id, name, email, password, created_at, updated_at FROM users WHERE name=$1"
	//}

	row, err := r.db.DB().ContextQuery(ctx, query, id)
	if err != nil {
		return &model.User{}, err
	}
	var user model.User
	for row.Next() {
		row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	}
	return &user, nil
}

func (r *repos) Create(ctx context.Context, user *model.UserInfo) (int, error) {
	if err := r.db.DB().Ping(ctx); err != nil {
		return 0, err
	}

	query := db.Query{
		Name:     "repository_create",
		QueryStr: "INSERT INTO users (name, email, password, roles) VALUES ($1, $2, $3, $4) RETURNING id",
	}

	row, err := r.db.DB().ContextQuery(ctx, query, user.Name, user.Email, user.Password, user.Role)
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
		Name:     "repository_delete",
		QueryStr: "DELETE FROM users WHERE id=$1",
	}
	_, err := r.db.DB().ContextExec(ctx, query, id)
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
		Name:     "repository_update",
		QueryStr: "UPDATE users SET name = $1, email = $2, updated_at = now() WHERE id=$3",
	}
	_, err := r.db.DB().ContextExec(ctx, query, user.Name, user.Email, user.Id)
	if err != nil {
		return err
	}
	return nil
}
