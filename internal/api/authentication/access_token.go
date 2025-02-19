package authentication

import (
	"context"
	auth "github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
)

func (i *ImplementationAuthentication) GetAccessToken(ctx context.Context, req *auth.GetAccessTokenRequest) (*auth.GetAccessTokenResponse, error) {
	accessToken, err := i.service.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}
	return &auth.GetAccessTokenResponse{AccessToken: accessToken}, nil
}
