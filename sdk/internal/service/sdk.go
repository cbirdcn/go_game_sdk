package service

import (
	"context"
	"fmt"
	v1 "sdk/api/sdk/v1"
	"sdk/internal/biz"
)

// service层，处理api请求
type SdkService struct {
	v1.UnimplementedSdkServer

	uc *biz.SdkUsecase
}

// NewSdkService：通过useCase实例化Service
func NewSdkService(uc *biz.SdkUsecase) *SdkService {
	return &SdkService{uc: uc}
}

// 处理Http/Grpc服务请求，从api/.../pb.go拿Req返回Reply。
// 利用req参数组装biz参数后交给useCase实例操作数据库，（如果有）返回符合格式需求的数据
func (s *SdkService) InitSdk(ctx context.Context, req *v1.InitSdkReq) (*v1.InitSdkReply, error) {
	packageInfo, err := s.uc.GetPackageData(ctx, &biz.InitSdkReq{
		AppId: req.AppId,
		Data: biz.InitSdkReqDataType{
			Udid:    req.Udid,
			Channel: req.Channel,
		},
		Sign: req.Sign,
	})

	// test data: hmset PACKAGE_INFO_CHANNEL_100 adId 1000001 channelId 100001 sonChannel 11000001 channelGroup 11 isBanReg false isBanPay false
	fmt.Println(packageInfo)

	return &v1.InitSdkReply{
		Code: 1,
		Msg: "ok",
	}, err
}