package data

// redis操作方法

import (
	"context"
	"sdk/internal/biz"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// Key and Prefix...

func packageInfoKeyPrefix() string{
	return "PACKAGE_INFO_CHANNEL_"
}

func gameInfoKeyPrefix() string{
	return "GAME_INFO_"
}

func banUserKeyPrefix() string{
	return "BAN_USER_"
}

func banImeiKeyPrefix() string{
	return "BAN_IMEI_"
}

func packageInfoBidKeyPrefix() string{
	return "PACKAGE_INFO_BID_"
}

func userInfoKeyPrefix() string{
	return "USER_INDEX_"
}

// function
// test data: hmset PACKAGE_INFO_CHANNEL_100 adId 1000001 channelId 100001 sonChannel 11000001 channelGroup 11 isBanReg false isBanPay false
func hotGetPackageInfo(ctx context.Context, rdb *redis.Client, channel uint32) (*biz.PackageInfo, error) {
	// HMget、HGetAll、MGet可以通过Scan扫入Struct
	var info biz.PackageInfo
	err := rdb.HMGet(ctx, packageInfoKeyPrefix() + strconv.Itoa(int(channel)),
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
	err := rdb.HMGet(ctx, gameInfoKeyPrefix() + strconv.Itoa(int(appId)),
		"gameId",
		"cpLoginKey").Scan(&info)
	return &info, err
}

func hotSetGameInfo(ctx context.Context, rdb *redis.Client, appId uint32, info *biz.Game) (bool) {
	gameInfo := make(map[string]interface{})
	gameInfo["gameId"] = info.GameId
	gameInfo["cpLoginKey"] = info.CpLoginKey
	res := rdb.HMSet(ctx, gameInfoKeyPrefix() + strconv.Itoa(int(appId)),gameInfo).Val()
	return res
}

func hotCheckUsernameInBanList(ctx context.Context, rdb *redis.Client, username string) (bool) {
	res := rdb.Exists(ctx, banUserKeyPrefix() + username).Val()
	if res == 1 {
		return true
	}
	return false
}

func hotCheckImeiInBanList(ctx context.Context, rdb *redis.Client, imei string) (bool) {
	res := rdb.Exists(ctx, banImeiKeyPrefix() + imei).Val()
	if res == 1 {
		return true
	}
	return false
}

func hotGetPackageInfoBidBanReg(ctx context.Context, rdb *redis.Client, channel uint32) (string) {
	channelString := strconv.FormatUint(uint64(channel), 10)
	return rdb.HGet(ctx, packageInfoBidKeyPrefix() + channelString, "isBanReg").Val()
}

func hotGetUserUidByUsername(ctx context.Context, rdb *redis.Client, username string) (string) {
	return rdb.HGet(ctx, userInfoKeyPrefix() + username, "uId").Val()
}

func hotGetPackageInfoBid(ctx context.Context, rdb *redis.Client, channel uint32) (map[string]string) {
	channelString := strconv.FormatUint(uint64(channel), 10)
	return rdb.HGetAll(ctx, packageInfoBidKeyPrefix() + channelString).Val()
}