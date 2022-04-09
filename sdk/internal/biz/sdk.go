package biz

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	v1 "sdk/api/sdk/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrorSetActiveRecordFail = errors.InternalServer(v1.ErrorReason_SetActiveRecordFail.String(), "set active record fail")
)

///////数据结构，在多个文件使用///////

// 将Req参数转成可以在多个模块间使用的不依赖Req参数的结构
type InitSdkReq struct {
	AppId uint32
	Data  InitSdkReqDataType
	Sign  string
}

type InitSdkReqDataType struct {
	Udid    string
	Channel uint32
}

// 模型，和db表名一致
type ActiveRecord struct {
	AId        uint32 `gorm:"primaryKey;column:aId"`
	BId        uint32 `gorm:"column:bId"`
	AdId       uint32 `gorm:"column:adId"`
	ChannelId  uint32 `gorm:"column:channelId"`
	SonChannel uint32 `gorm:"column:sonChannel"`
	ActiveTime uint32 `gorm:"column:activeTime"`
	GameId     uint32 `gorm:"column:gameId"`
	DownIp     uint32 `gorm:"column:downIp"`
	ImeiCode   string `gorm:"column:imeiCode"`
	IsNewImei  uint8  `gorm:"column:isNewImei"`
}

// 数据模型，给data用操作db
// redis tag 提供给data/redis.go做hmget scan 成 struct
type PackageInfo struct {
	BId        uint32 `redis:"bid"`
	AdId       uint32 `redis:"adId"`
	ChannelId  uint32 `redis:"channelId"`
	SonChannel uint32 `redis:"sonChannel"`
	// ChannelGroup uint32 `redis:"channelGroup"`
	// IsBanReg     bool   `redis:"isBanReg"`
	// IsBanPay     bool   `redis:"isBanPay"`
}

// repo业务逻辑接口：在data中实现
type SdkRepo interface {
	GetPackageInfo(ctx context.Context, request *InitSdkReq) (*PackageInfo, error)
	SetActiveRecord(ctx context.Context, request *InitSdkReq) (bool, error)
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

// TODO:
func HttpBuildQuery(request *InitSdkReq) (param_str string) {
	return ""
}

// TODO:
func (uc *SdkUsecase) CheckSign(ctx context.Context, request *InitSdkReq) (bool, error) {
	// 请求签名方式：appid + service_name + data字符串 + login_key
	// service_name可以是sdk.game.initsdk或其他
	// login_key是后台给每个游戏配置的唯一秘钥
	// data字符串结构：url.QueryEscape(param1=a&param2=b)
	return true, nil
}

func (uc *SdkUsecase) GetPackageData(ctx context.Context, request *InitSdkReq) (*PackageInfo, error) {
	return uc.repo.GetPackageInfo(ctx, request)
}

func (uc *SdkUsecase) SetActiveRecord(ctx context.Context, request *InitSdkReq) (bool, error) {
	return uc.repo.SetActiveRecord(ctx, request)
}
