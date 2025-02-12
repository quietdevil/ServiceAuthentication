package user

import (
	"context"
	"serviceauth/internal/model"
	informationmethod "serviceauth/internal/repository/Information_method"
	r "serviceauth/internal/repository/logs/model"
)

func (s *serviceUser) Create(ctx context.Context, user *model.UserInfo) (int, error) {
	var id int
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		l := r.Logs{Name: user.Name, Description: "Create User service layer success", MethodName: informationmethod.MethodCreate}
		err := s.logs.Create(ctx, l)
		if err != nil {
			return err
		}
		idU, err := s.userRepository.Create(ctx, user)
		id = idU
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *serviceUser) Delete(ctx context.Context, id int) error {
	if err := s.userRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s *serviceUser) Get(ctx context.Context, id int) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}

func (s *serviceUser) Update(ctx context.Context, user *model.UserUpdate) error {
	if err := s.userRepository.Update(ctx, user); err != nil {
		return err
	}
	return nil
}
