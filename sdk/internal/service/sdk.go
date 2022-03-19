package service

import (
	"context"

	v1 "sdk/api/sdk/v1"
	"sdk/internal/biz"
)

// SdkService is a greeter service.
type SdkService struct {
	v1.UnimplementedSdkServer

	uc *biz.SdkUsecase
}

// NewSdkService new a greeter service.
func NewSdkService(uc *biz.SdkUsecase) *SdkService {
	return &SdkService{uc: uc}
}

// SayHello implements sdk.SdkServer.
func (s *SdkService) SayHello(ctx context.Context, in *v1.CheckEnterRequest) (*v1.CheckEnterReply, error) {
	//g, err := s.uc.CreateSdk(ctx, &biz.Sdk{Hello: in.Name})
	//if err != nil {
	//	return nil, err
	//}
	return &v1.CheckEnterReply{}, nil
}
