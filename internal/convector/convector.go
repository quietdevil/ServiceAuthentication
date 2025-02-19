package convector

import (
	"github.com/quietdevil/ServiceAuthentication/internal/model"
	desc "github.com/quietdevil/ServiceAuthentication/pkg/auth_user_v1"
	auth "github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromGrpsIntoModel(req *desc.CreateRequest) *model.UserInfo {
	return &model.UserInfo{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role.String(),
	}
}

func FromModelIntoGrpc(model *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if model.UpdatedAt.Valid {
		updatedAt = timestamppb.New(model.UpdatedAt.Time)
	}

	v, _ := desc.Enum_value[model.Role]

	return &desc.GetResponse{
		Id:        int64(model.Id),
		Name:      model.UserInfo.Name,
		Email:     model.UserInfo.Email,
		Role:      desc.Enum(v),
		CreatedAt: timestamppb.New(model.CreatedAt),
		UpdatedAt: updatedAt,
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
