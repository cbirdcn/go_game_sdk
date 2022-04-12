# Go分布式手游平台Server SDK

## Feature

- Kratos作为基础框架

## 进度

- 初始化项目
- 完成App激活功能：initSdk

## Tutorial

1. 按照 [Kratos官方初始化](https://go-kratos.dev/docs/getting-started/start/) 方式安装kratos/v2
2. go mod tidy 整理mod
3. go generate ./...自动生成，如果缺少wire，按照提示安装
4. kratos run 运行可执行文件，直到启动grpc和http服务
5. 修改proto文件后，可以用 [工具](https://go-kratos.dev/docs/getting-started/usage/) 自动生成pb.go
6. curl -X POST http://127.0.0.1:8000/sdk/check_enter
7. 如果wire_gen.go生成失败或无法通过generate修正，要手动修改文件

## 感谢

- [kratos](https://github.com/go-kratos/kratos)
- [google protobuf](https://developers.google.com/protocol-buffers)
- [go-redis](github.com/go-redis/redis/v8)
- [consul](https://github.com/hashicorp/consul)
- [wire](github.com/google/wire)
- [gorm](https://github.com/go-gorm/gorm)
- [ent](https://entgo.io/docs/getting-started)

## 功能说明

- 验签
验签处理在middleware中
签名方式：appid + service_name + data字符串 + login_key
service_name可以是sdk.game.initsdk或其他
data字符串结构：url.QueryEscape(param1=a&param2=b)，其中paramN是data内的所有参数
login_key是后台给每个游戏配置的唯一秘钥
当cpLoginKey=b5d6ad09709d5eac958e8fb8ad511c76时，数据和签名如下

```json
{
  "service": "sdk.game.initsdk",
  "appId": 1000000,
  "data": {
    "udid": "12-34-56-78-9100",
    "channel": 1000001
  },
  "sign": "ad4522e0fcad8c671d679e0e19c75968"
}
```

- sdk.game.initSdk
下载游戏后首次打开游戏，首先初始化sdk
将udid、channel、cploginkey、appid，以及签名sign提交到服务器
根据参数到redis查渠道包信息，判断是否新设备（7天内没有激活记录）
写入激活记录到DB中(myx_log.active_records表)

```go
curl --location --request POST 'http://127.0.0.1:8000/sdk/api/v1/init_sdk' 
--header 'Content-Type: application/json' 
--data-raw '{
    "service": "sdk.game.initsdk",
    "appId": 1000000,
    "data": {
        "udid": "12-34-56-78-9100",
        "channel": 1000001
    },
    "sign": "ad4522e0fcad8c671d679e0e19c75968"
}'
```

- sdk.user.reg
```go
{
  "service": "sdk.user.reg",
  "appId": 1000000,
  "data": {
    "username": "Hello",
    "passwd": "Hello",
    "udid": "12-34-56-78-9100",
    "channel": 1000001
  },
  "sign": "a58154b890baf0022b36f089263b9312"
}
```