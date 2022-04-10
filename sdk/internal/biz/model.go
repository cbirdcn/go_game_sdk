package biz
///////数据结构，在多个文件使用///////

// 将Req参数转成可以在多个模块间使用的不依赖Req参数的结构
type InitSdkReq struct {
	Service string
	AppId   uint32
	Data    InitSdkReqDataType
	Sign    string
}

type InitSdkReqDataType struct {
	Udid    string
	Channel uint32
}

// 模型，和db表名一致

// 激活日志记录表
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

// 包信息
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

// game info
// TODO: 需要继续补充数据
type Game struct {
	GameId uint32 `gorm:"column:gameId"`
	CpLoginKey string `gorm:"column:cpLoginKey"`
}

type GameInfo struct {
	GameId uint32 `redis:"gameId"`
	CpLoginKey string `redis:"cpLoginKey"`
}