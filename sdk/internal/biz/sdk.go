package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)


// repo业务逻辑接口：在data中实现
type SdkRepo interface {
	GetPackageInfo(ctx context.Context, request *InitSdkReq) (*PackageInfo, error)
	SetActiveRecord(ctx context.Context, request *InitSdkReq) (bool, error)
	GetGameInfo(ctx context.Context, appId uint32) (*GameInfo, error)
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

func (uc *SdkUsecase) GetPackageData(ctx context.Context, request *InitSdkReq) (*PackageInfo, error) {
	return uc.repo.GetPackageInfo(ctx, request)
}

func (uc *SdkUsecase) SetActiveRecord(ctx context.Context, request *InitSdkReq) (bool, error) {
	return uc.repo.SetActiveRecord(ctx, request)
}

func (uc *SdkUsecase) GetGameData(ctx context.Context, appId uint32) (*GameInfo, error) {
	return uc.repo.GetGameInfo(ctx, appId)
}
