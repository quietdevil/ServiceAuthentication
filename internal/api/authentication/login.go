package authentication

import (
	"context"
	"github.com/quietdevil/ServiceAuthentication/internal/convector"
	auth "github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
)

func (i *ImplementationAuthentication) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	token, err := i.service.Login(ctx, convector.FromGprcLogIntoUserLogin(req))
	if err != nil {
		return nil, err
	}
	return &auth.LoginResponse{
		RefreshToken: token,
	}, nil

}
