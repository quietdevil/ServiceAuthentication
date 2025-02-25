package access

import (
	"context"
	"github.com/quietdevil/ServiceAuthentication/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"strings"
)

func (a *AccessService) Check(ctx context.Context, endpoint string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.InvalidArgument, "Missing Metadata")
	}
	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) != 1 {
		return status.Error(codes.InvalidArgument, "Missing Authorization Token+Header")
	}
	if !strings.HasPrefix(authHeader[0], a.configAuth.AuthenticationPrefix()) {
		return status.Error(codes.InvalidArgument, "Invalid Authorization Header")
	}

	token := strings.TrimPrefix(authHeader[0], a.configAuth.AuthenticationPrefix())

	claims, err := utils.VerifyToken(token[1:], []byte(a.configAuth.AccessSecretKey()))

	if err != nil {
		return err
	}

	roleName, err := a.reposAccess.Role(ctx, endpoint)
	if err != nil {
		return err
	}

	if claims.Role != roleName {
		return status.Error(codes.PermissionDenied, "access denied")
	}

	return nil
}
