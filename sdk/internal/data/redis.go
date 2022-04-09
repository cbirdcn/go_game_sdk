package data

// redis操作方法

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sdk/internal/biz"
	"strconv"
)

// Key and Prefix...

func getPackageInfoKeyPrefix() string{
	return "PACKAGE_INFO_CHANNEL_"
}

// function
// test data: hmset PACKAGE_INFO_CHANNEL_100 adId 1000001 channelId 100001 sonChannel 11000001 channelGroup 11 isBanReg false isBanPay false
func getPackageInfo(ctx context.Context, rdb *redis.Client, channel uint32) (*biz.PackageInfo, error) {
	// HMget、HGetAll、MGet可以通过Scan扫入Struct
	var packageInfo biz.PackageInfo
	err := rdb.HMGet(ctx, getPackageInfoKeyPrefix() + strconv.Itoa(int(channel)),
		"bid",
		"adId",
		"channelId",
		"sonChannel",
		"channelGroup",
		"isBanReg",
		"isBanPay").Scan(&packageInfo)
	return &packageInfo, err
}
