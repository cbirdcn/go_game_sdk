package data

import (
	"context"

	"sdk/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

// repo层：用于实现biz层中useCase抽象的方法
type sdkRepo struct {
	data *Data
	log  *log.Helper
}

// useCase方法的实现：直接与DB数据打交道
func (r *sdkRepo) GetPackageInfo(ctx context.Context, request *biz.InitSdkReq) (*biz.PackageInfoType, error) {
	return getPackageInfo(ctx, r.data.rdb, request.Data.Channel)
}

// NewSdkRepo .
func NewSdkRepo(data *Data, logger log.Logger) biz.SdkRepo {
	return &sdkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}