package data

import (
	"context"

	"sdk/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type sdkRepo struct {
	data *Data
	log  *log.Helper
}

// NewSdkRepo .
func NewSdkRepo(data *Data, logger log.Logger) biz.SdkRepo {
	return &sdkRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *sdkRepo) Save(ctx context.Context, g *biz.Sdk) (*biz.Sdk, error) {
	return g, nil
}

func (r *sdkRepo) Update(ctx context.Context, g *biz.Sdk) (*biz.Sdk, error) {
	return g, nil
}

func (r *sdkRepo) FindByID(context.Context, int64) (*biz.Sdk, error) {
	return nil, nil
}

func (r *sdkRepo) ListByHello(context.Context, string) ([]*biz.Sdk, error) {
	return nil, nil
}

func (r *sdkRepo) ListAll(context.Context) ([]*biz.Sdk, error) {
	return nil, nil
}
