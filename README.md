# Go分布式手游平台Server SDK

## Feature
- Kratos作为基础框架

## 进度
- 初始化项目

## Tutorial
1. 按照 [Kratos官方初始化](https://go-kratos.dev/docs/getting-started/start/) 方式安装kratos/v2
1. go mod tidy 整理mod 
1. go generate ./...自动生成，如果缺少wire，按照提示安装
1. kratos run 运行可执行文件，直到启动grpc和http服务
1. 修改proto文件后，可以用 [工具](https://go-kratos.dev/docs/getting-started/usage/) 自动生成pb.go
1. curl -X POST http://127.0.0.1:8000/sdk/check_enter
1. 如果wire_gen.go生成失败或无法通过generate修正，要手动修改文件

## 感谢
- [kratos](https://github.com/go-kratos/kratos)
- [google protobuf](https://developers.google.com/protocol-buffers)
- [go-redis](github.com/go-redis/redis/v8)
- [consul](https://github.com/hashicorp/consul)
- [wire](github.com/google/wire)