package logic

import (
	"context"

	"rpc/internal/pb"
	"rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 展示用户列表
func (l *ListUsersLogic) ListUsers(in *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.ListUsersResponse{}, nil
}
