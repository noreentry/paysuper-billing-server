package service

import (
	"context"
	"github.com/paysuper/paysuper-proto/go/billingpb"
)

func (s *Service) FindByZipCode(
	ctx context.Context,
	req *billingpb.FindByZipCodeRequest,
	rsp *billingpb.FindByZipCodeResponse,
) error {
	if req.Country != CountryCodeUSA {
		return nil
	}

	count, err := s.zipCodeRepository.CountByZip(ctx, req.Zip, req.Country)

	if err != nil {
		return orderErrorUnknown
	}

	if count <= 0 {
		return nil
	}

	data, err := s.zipCodeRepository.FindByZipAndCountry(ctx, req.Zip, req.Country, req.Offset, req.Limit)

	if err != nil {
		return orderErrorUnknown
	}

	rsp.Count = int32(count)
	rsp.Items = data

	return nil
}
