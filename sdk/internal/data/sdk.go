package data

import (
	"context"
	"errors"

	// "fmt"
	"time"

	"sdk/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	// "github.com/go-redis/redis/v8"
	// "github.com/thinkeridea/go-extend"
	// "github.com/thinkeridea/go-extend/exnet"
)

// repo层：用于实现biz层中useCase抽象的方法
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

// useCase方法的实现：直接与DB、redis数据打交道

// 获取redis中包信息
func (r *sdkRepo) GetPackageInfo(ctx context.Context, request *biz.InitSdkReq) (*biz.PackageInfo, error) {
	return getPackageInfo(ctx, r.data.rdb, request.Data.Channel)
}

// deprecated：ent的实现方式不能支持自定义主键名称。改用gorm
// func (r *sdkRepo) SetActiveRecord(ctx context.Context, request *biz.InitSdkReq) (bool, error) {
// 	// 暂时放弃自定义时区
// 	nowTime := time.Now()
// 	beforeTime := nowTime.AddDate(0, 0, -7).Unix()
// 	_, err := r.data.db.ActiveRecord.
// 		Query().
// 		Where(activerecord.ImeiCode(request.Data.Udid)).
// 		Where(activerecord.ActiveTimeGT(uint32(beforeTime))).First(ctx)
// 	isnewImei := 1
// 	if err == nil {
// 		isnewImei = 0
// 	}
// 	packageInfo, err := r.GetPackageInfo(ctx, request)
// 	if err != nil {
// 		return false, fmt.Errorf("failed get package info: %w", err)
// 	}
// 	fakeIp := uint64(10000) // todo

// 	_, err = r.data.db.ActiveRecord.Create().
// 		SetID(1).
// 		SetActiveTime(uint32(nowTime.Unix())).
// 		SetAdId(packageInfo.AdId).
// 		SetBId(packageInfo.BId).
// 		SetChannelId(packageInfo.ChannelId).
// 		SetSonChannel(packageInfo.SonChannel).
// 		SetGameId(request.AppId).
// 		SetDownIp(fakeIp).
// 		SetImeiCode(request.Data.Udid).
// 		SetIsnewImei(uint8(isnewImei)).
// 		Save(ctx)

// 	if err != nil {
// 		return false, fmt.Errorf("failed save active record: %w", err)
// 	}

// 	return true, nil
// }

// 激活log写入db
func (r *sdkRepo) SetActiveRecord(ctx context.Context, request *biz.InitSdkReq) (bool, error) {
	var activeRecord biz.ActiveRecord
	nowTime := time.Now()
	beforeTime := nowTime.AddDate(0, 0, -7).Unix()
	res := r.data.db.Where("imeiCode = ?", request.Data.Udid).Where("activeTime >= ?", beforeTime).Limit(1).Find(&activeRecord)

	if request.Data.Udid == "" {
		return false,errors.New("device not found.")
	}

	packageInfo, err := r.GetPackageInfo(ctx, request)
	if err != nil {
		return false, err
	}
	// clientIp, err := exnet.IPString2Long(exnet.ClientIP(request)) // 需要从ctx获取header。

	newActiveRecord := biz.ActiveRecord{
		BId          : packageInfo.BId,
		AdId         : packageInfo.AdId,
		ChannelId    : packageInfo.ChannelId,
		SonChannel   : packageInfo.SonChannel,
		ActiveTime	 : uint32(nowTime.Unix()),
		GameId		 : request.AppId,
		DownIp     	 : 123, // TODO: 需要获取客户端ip并转成uint
		ImeiCode     : request.Data.Udid,
		IsNewImei    : 0,
	}
	if res.RowsAffected == 0 {
		newActiveRecord.IsNewImei = 1
	}

	res = r.data.db.Create(&newActiveRecord)
	if res.RowsAffected == 1 {
		return true,nil
	}
	return false,res.Error
}
