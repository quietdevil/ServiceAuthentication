package access

import (
	"context"
	"github.com/quietdevil/ServiceAuthentication/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *ImplementationAccess) Check(ctx context.Context, req *access_v1.CheckRequest) (*emptypb.Empty, error) {

	err := i.serviceLayer.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil

}
