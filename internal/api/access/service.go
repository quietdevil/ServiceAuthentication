package access

import (
	"serviceauth/internal/service"
	"serviceauth/pkg/access_v1"
)

type ImplementationAccess struct {
	access_v1.UnimplementedAccessV1Server
	serviceLayer service.AccessService
}

func NewImplementationAccess(serviceLayer service.AccessService) *ImplementationAccess {
	return &ImplementationAccess{serviceLayer: serviceLayer}
}
