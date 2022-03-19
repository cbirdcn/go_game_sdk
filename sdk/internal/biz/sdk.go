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

// Sdk is a Sdk model.
type Sdk struct {
	Hello string
}

// SdkRepo is a Greater repo.
type SdkRepo interface {
	Save(context.Context, *Sdk) (*Sdk, error)
	Update(context.Context, *Sdk) (*Sdk, error)
	FindByID(context.Context, int64) (*Sdk, error)
	ListByHello(context.Context, string) ([]*Sdk, error)
	ListAll(context.Context) ([]*Sdk, error)
}

// SdkUsecase is a Sdk usecase.
type SdkUsecase struct {
	repo SdkRepo
	log  *log.Helper
}

// NewSdkUsecase new a Sdk usecase.
func NewSdkUsecase(repo SdkRepo, logger log.Logger) *SdkUsecase {
	return &SdkUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateSdk creates a Sdk, and returns the new Sdk.
func (uc *SdkUsecase) CreateSdk(ctx context.Context, g *Sdk) (*Sdk, error) {
	uc.log.WithContext(ctx).Infof("CreateSdk: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
