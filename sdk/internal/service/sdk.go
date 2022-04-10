package service

import (
	"context"
	// "github.com/go-kratos/kratos/v2/errors"
	v1 "sdk/api/sdk/v1"
	"sdk/internal/biz"
)

// service层，处理api请求
type SdkService struct {
	v1.UnimplementedSdkServer

	uc *biz.SdkUsecase
}

// NewSdkService：通过useCase实例化Service
func NewSdkService(uc *biz.SdkUsecase) *SdkService {
	return &SdkService{uc: uc}
}

// 处理Http/Grpc服务请求，从api/.../pb.go拿Req返回Reply。
// 利用req参数组装biz参数后交给useCase实例操作数据库，（如果有）返回符合格式需求的数据
func (s *SdkService) InitSdk(ctx context.Context, req *v1.InitSdkReq) (*v1.InitSdkReply, error) {
	// 中间件验证签名，业务逻辑中不需要处理

	// 格式化输入参数
	initSdkReq := &biz.InitSdkReq{
		Service: req.Service,
		AppId:   req.AppId,
		Data: biz.InitSdkReqDataType{
			Udid:    req.Data.Udid,
			Channel: req.Data.Channel,
		},
		Sign: req.Sign,
	}
	// 调用usecase的方法，写入格式化的数据到DB。
	// 具体逻辑由useCase声明并交给data实现，service这里只处理调用和返回。
	res, err := s.uc.SetActiveRecord(ctx, initSdkReq)

	// 避免使用err判断，因为err可能是gorm.Err等类型，比如NotFoundErr
	if res {
		return &v1.InitSdkReply{
			Code: 0,
			Msg:  "ok",
		}, err
	} else {
		return &v1.InitSdkReply{
			Code: 1,
			Msg:  "set active record failed",
		}, biz.ErrorSetActiveRecordFail
	}

}
