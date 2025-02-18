package access

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"serviceauth/pkg/access_v1"
)

func (i *ImplementationAccess) Check(ctx context.Context, req *access_v1.CheckRequest) (*emptypb.Empty, error) {

	err := i.serviceLayer.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil

}
