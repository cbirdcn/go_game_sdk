package data

// redis操作方法

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"sdk/internal/biz"
	"strconv"
)

// Key and Prefix...

func getPackageInfoKeyPrefix() string{
	return "PACKAGE_INFO_CHANNEL_"
}

// function

func getPackageInfo(ctx context.Context, rdb *redis.Client, channel uint32) (*biz.PackageInfoType, error) {
	// HMget、HGetAll、MGet可以通过Scan扫入Struct
	var packageInfo biz.PackageInfoType
	err := rdb.HMGet(ctx, getPackageInfoKeyPrefix() + strconv.Itoa(int(channel)),
		"adId",
		"channelId",
		"sonChannel",
		"channelGroup",
		"isBanReg",
		"isBanPay").Scan(&packageInfo)
	return &packageInfo, err
}
