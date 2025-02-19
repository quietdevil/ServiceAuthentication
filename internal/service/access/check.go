package access

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"serviceauth/internal/service/authentication"
	"serviceauth/internal/utils"
	"strings"
)

const authPrefix = "Bearer"

func (a *AccessService) Check(ctx context.Context, endpoint string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.InvalidArgument, "Missing Metadata")
	}
	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) != 1 {
		return status.Error(codes.InvalidArgument, "Missing Authorization Token+Header")
	}
	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return status.Error(codes.InvalidArgument, "Invalid Authorization Header")
	}

	token := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(token[1:], []byte(authentication.AccessSecretKey))

	if err != nil {
		return err
	}
	//todo Обращение в бд за правами

	roleName, err := a.reposAccess.Role(ctx, endpoint)
	if err != nil {
		return err
	}

	if claims.Role != roleName {
		return status.Error(codes.PermissionDenied, "access denied")
	}

	return nil
}

func accessHandle(context.Context) map[string]int32 {
	accessHandleMap := make(map[string]int32)
	accessHandleMap["/auth_v1.AuthenticationUserV1/Get"] = 0
	return accessHandleMap
}
