package user

import (
	"context"
	"fmt"
	"serviceauth/internal/convector"
	desc "serviceauth/pkg/auth_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, int(req.GetId()))
	if err != nil {
		return &desc.GetResponse{}, err
	}
	return convector.FromModelIntoGrpc(user), nil

}

func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.userService.Create(ctx, convector.FromGrpsIntoModel(req))
	if err != nil {
		fmt.Println(err)
		return &desc.CreateResponse{}, err

	}
	return &desc.CreateResponse{Id: int64(id)}, nil
}

func (s *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	if err := s.userService.Update(ctx, convector.FromGrpsIntoModelUpdate(req)); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if err := s.userService.Delete(ctx, 3); err != nil {
		return &emptypb.Empty{}, nil
	}
	return &emptypb.Empty{}, nil
}
