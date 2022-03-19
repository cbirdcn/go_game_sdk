package biz

import (
	"context"

	v1 "sdk/api/sdk/v1"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)


// 数据模型，在多个文件使用
type InitSdkReq struct {
	AppId uint32
	Data InitSdkReqDataType
	Sign string
}

type InitSdkReqDataType struct {
	Udid string
	Channel uint32
}

// redis tag 提供给data/redis.go做hmget scan 成 struct
type PackageInfoType struct {
	AdId uint32 `redis:"adId"`
	ChannelId uint32 `redis:"channelId"`
	SonChannel uint32 `redis:"sonChannel"`
	ChannelGroup uint32 `redis:"channelGroup"`
	IsBanReg bool `redis:"isBanReg"`
	IsBanPay bool `redis:"isBanPay"`
}

// repo业务逻辑接口：在data中实现
type SdkRepo interface {
	GetPackageInfo(ctx context.Context, request *InitSdkReq) (*PackageInfoType, error)
}

// useCase，在service中注入后，可以调用useCase的方法
type SdkUsecase struct {
	repo SdkRepo
	log  *log.Helper
}

// useCase创建实例
func NewSdkUsecase(repo SdkRepo, logger log.Logger) *SdkUsecase {
	return &SdkUsecase{repo: repo, log: log.NewHelper(logger)}
}

// useCase的方法：给service调用。内部调用repo的方法操作repo
func (uc *SdkUsecase) GetPackageData(ctx context.Context, request *InitSdkReq) (*PackageInfoType, error) {
	return uc.repo.GetPackageInfo(ctx, request)
}
