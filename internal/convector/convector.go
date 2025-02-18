package convector

import (
	"serviceauth/internal/model"
	desc "serviceauth/pkg/auth_user_v1"
	auth "serviceauth/pkg/auth_v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromGrpsIntoModel(req *desc.CreateRequest) *model.UserInfo {
	return &model.UserInfo{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}

func FromModelIntoGrpc(model *model.User) *desc.GetResponse {
	return &desc.GetResponse{
		Id:        int64(model.Id),
		Name:      model.UserInfo.Name,
		Email:     model.UserInfo.Email,
		CreatedAt: timestamppb.New(model.Created_at),
		UpdatedAt: timestamppb.New(model.Updated_at),
	}
}

func FromGrpsIntoModelUpdate(req *desc.UpdateRequest) *model.UserUpdate {
	return &model.UserUpdate{
		Id:    int(req.GetId()),
		Name:  req.GetName().Value,
		Email: req.GetEmail().Value,
	}
}

func FromGprcLogIntoUserLogin(req *auth.LoginRequest) *model.UserLogin {
	return &model.UserLogin{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}
}
