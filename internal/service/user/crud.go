package user

import (
	"context"
	"serviceauth/internal/model"
)

func (s *serviceUser) Create(ctx context.Context, user *model.UserInfo) (int, error) {
	id, err := s.userRepository.Create(ctx, user)
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
