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

type RegReq struct {
	Service string
	AppId   uint32
	Data    RegReqDataType
	Sign    string
}

type RegReqDataType struct {
	Username string
	Passwd   string
	Udid     string
	Channel  uint32
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

// 账号索引表
type UserIndex struct {
	UId           uint32 `gorm:"primaryKey;column:uId;autoIncrement"`
	UserAcount    string `gorm:"column:userAccount"`
	RegTime       uint32 `gorm:"column:regTime"`
	AdId          uint32 `gorm:"column:adId"`
	ChannelId     uint32 `gorm:"column:channelId"`
	SonChannel    uint32 `gorm:"column:sonChannel"`
	AgentId       uint32 `gorm:"column:agentId"`
	placeId       uint32 `gorm:"column:placeId"`
	ChannelGroup  uint32 `gorm:"column:channelGroup"`
	RegGameId     uint32 `gorm:"column:regGameId"`
	Udid          string `gorm:"column:udid"`
	Mobile        uint16 `gorm:"column:mobile"`
	Platform      uint32 `gorm:"column:platform;default:1"`
	PlatformMoney uint32 `gorm:"column:platformMoney"`
	Extension     string `gorm:"column:extension"`
}

// 账号注册记录表
type RegRecord struct {
	Id            uint32 `gorm:"column:id;autoIncrement"`
	userId        uint32 `gorm:"column:userId;unique"`
	UserAcount    string `gorm:"column:userAccount;unique"`
	Mobile        uint16 `gorm:"column:mobile;unique"`
	RegTime       uint32 `gorm:"column:regTime"`
	AdId          uint32 `gorm:"column:adId"`
	ChannelId     uint32 `gorm:"column:channelId"`
	SonChannel    uint32 `gorm:"column:sonChannel"`
	AgentId       uint32 `gorm:"column:agentId"`
	placeId       uint32 `gorm:"column:placeId"`
	ChannelGroup  uint32 `gorm:"column:channelGroup"`
	RegGameId     uint32 `gorm:"column:regGameId"`
	RegIp         uint32 `gorm:"column:regIp"`
	Province      string `gorm:"column:province;default:其他"`
	RegProvince   string `gorm:"column:regProvince"`
	RegImei       string `gorm:"column:regImei"`
	IsnewImei     string `gorm:"column:isnewImei;default:1"`
	Platform      uint32 `gorm:"column:platform;default:1"`
	PlatformMoney uint32 `gorm:"column:platformMoney"`
	Extension     string `gorm:"column:extension"`
}

// user_i表，需要临时指令表名，i为hash值0-254
type User struct {
	Id            uint32 `gorm:"column:id;autoIncrement"`
	UId           uint32 `gorm:"column:uId;unique"`
	UserAcount    string `gorm:"column:userAccount"`
	UserPass      string `gorm:"column:userPass"`
	PayPassword   string `gorm:"column:payPassword"`
	PlatformMoney uint32 `gorm:"column:platformMoney"`
	UserNickName  string `gorm:"column:userNickName"`
	UserProfile   uint8  `gorm:"column:userProfile"`
	UserBirthday  string `gorm:"column:userBirthday"`
	Mobile        uint16 `gorm:"column:mobile"`
	Email         string `gorm:"column:email"`
	IdCard        string `gorm:"column:idCard"`
	TrueName      string `gorm:"column:trueName"`
	UserSex       uint8  `gorm:"column:userSex"`
	Score         uint32 `gorm:"column:score"`
	RegTime       uint32 `gorm:"column:regTime"`
	RegIp         uint32 `gorm:"column:regIp"`
	RegImei       string `gorm:"column:regImei`
	LastLoginTime uint32 `gorm:"column:lastLoginTi"`
	LastLoginIp   uint32 `gorm:"column:lastLoginIp"`
	ChannelId     uint32 `gorm:"column:channelId"`
	SonChannel    uint32 `gorm:"column:sonChannel"`
	AgentId       uint32 `gorm:"column:agentId"`
	PlaceId       uint32 `gorm:"column:placeId"`
	RegGameId     uint32 `gorm:"column:regGameId"`
	State         uint8  `gorm:"column:state"`
	Extension     string `gorm:"column:extension"`
	Pwd           string `gorm:"column:pwd"`
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
	GameId     uint32 `gorm:"column:gameId"`
	CpLoginKey string `gorm:"column:cpLoginKey"`
}

type GameInfo struct {
	GameId     uint32 `redis:"gameId"`
	CpLoginKey string `redis:"cpLoginKey"`
}
