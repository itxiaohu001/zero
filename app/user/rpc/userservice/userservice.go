// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2
// Source: user.proto

package userservice

import (
	"context"

	"rpc/internal/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserDetailsRequest  = pb.GetUserDetailsRequest
	GetUserDetailsResponse = pb.GetUserDetailsResponse
	ListUsersRequest       = pb.ListUsersRequest
	ListUsersResponse      = pb.ListUsersResponse
	LoginRequest           = pb.LoginRequest
	LoginResponse          = pb.LoginResponse
	RegisterRequest        = pb.RegisterRequest
	RegisterResponse       = pb.RegisterResponse
	User                   = pb.User
	UserInfoRequest        = pb.UserInfoRequest
	UserInfoResponse       = pb.UserInfoResponse

	UserService interface {
		// 注册
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		// 登录
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		// 获取用户信息
		GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
		// 展示用户列表
		ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error)
		// 获取用户详细信息
		GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

// 注册
func (m *defaultUserService) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 登录
func (m *defaultUserService) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 获取用户信息
func (m *defaultUserService) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

// 展示用户列表
func (m *defaultUserService) ListUsers(ctx context.Context, in *ListUsersRequest, opts ...grpc.CallOption) (*ListUsersResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.ListUsers(ctx, in, opts...)
}

// 获取用户详细信息
func (m *defaultUserService) GetUserDetails(ctx context.Context, in *GetUserDetailsRequest, opts ...grpc.CallOption) (*GetUserDetailsResponse, error) {
	client := pb.NewUserServiceClient(m.cli.Conn())
	return client.GetUserDetails(ctx, in, opts...)
}
