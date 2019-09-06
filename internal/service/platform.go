package service

import (
	"context"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-billing-server/pkg/proto/grpc"
)

func (s *Service) GetPlatforms(ctx context.Context, req *grpc.ListPlatformsRequest, rsp *grpc.ListPlatformsResponse) error {
	platforms := []*grpc.Platform{}
	rsp.Status = pkg.ResponseStatusOk
	var i int32
	for _, pl := range availablePlatforms {
		if i < req.Offset {
			i++
			continue
		}

		platforms = append(platforms, pl)
		if int32(len(platforms)) >= req.Limit {
			break
		}
		i++
	}
	rsp.Platforms = platforms

	return nil
}
