package logic

import (
	"context"

	"rpc/internal/pb"
	"rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDetailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDetailsLogic {
	return &GetUserDetailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户详细信息
func (l *GetUserDetailsLogic) GetUserDetails(in *pb.GetUserDetailsRequest) (*pb.GetUserDetailsResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserDetailsResponse{}, nil
}
