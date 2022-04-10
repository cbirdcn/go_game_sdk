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

func getGameInfoKeyPrefix() string{
	return "GAME_INFO_"
}

// function
// test data: hmset PACKAGE_INFO_CHANNEL_100 adId 1000001 channelId 100001 sonChannel 11000001 channelGroup 11 isBanReg false isBanPay false
func hotGetPackageInfo(ctx context.Context, rdb *redis.Client, channel uint32) (*biz.PackageInfo, error) {
	// HMget、HGetAll、MGet可以通过Scan扫入Struct
	var info biz.PackageInfo
	err := rdb.HMGet(ctx, getPackageInfoKeyPrefix() + strconv.Itoa(int(channel)),
		"bid",
		"adId",
		"channelId",
		"sonChannel",
		"channelGroup",
		"isBanReg",
		"isBanPay").Scan(&info)
	return &info, err
}

func hotGetGameInfo(ctx context.Context, rdb *redis.Client, appId uint32) (*biz.GameInfo, error) {
	var info biz.GameInfo
	err := rdb.HMGet(ctx, getGameInfoKeyPrefix() + strconv.Itoa(int(appId)),
		"gameId",
		"cpLoginKey").Scan(&info)
	return &info, err
}

func hotSetGameInfo(ctx context.Context, rdb *redis.Client, appId uint32, info *biz.Game) (bool) {
	gameInfo := make(map[string]interface{})
	gameInfo["gameId"] = info.GameId
	gameInfo["cpLoginKey"] = info.CpLoginKey
	res := rdb.HMSet(ctx, getGameInfoKeyPrefix() + strconv.Itoa(int(appId)),gameInfo).Val()
	return res
}