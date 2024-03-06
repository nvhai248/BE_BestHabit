package usergrpcserver

import (
	"bestHabit/common"
	proto "bestHabit/generatedProto/proto/userservice"
	"bestHabit/modules/user/usermodel"
	"bestHabit/modules/user/userstorage"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCServer struct {
	db *sqlx.DB
	proto.UnimplementedUserServiceServer
}

func NewGRPCServer(db *sqlx.DB) *gRPCServer {
	return &gRPCServer{db: db}
}

func (r *gRPCServer) UserUpdateProfile(ctx context.Context, request *proto.UserUpdateProfileRequest) (*proto.UserUpdateProfileResponse, error) {
	fmt.Println("OK 1")

	storage := userstorage.NewSQLStore(r.db)

	fmt.Println("OK 1")
	err := storage.UpdateInfoById(ctx, &usermodel.UserUpdate{
		Name:  &request.Name,
		Phone: &request.Phone,
		Avatar: &common.Image{
			Url:       request.Avatar.Url,
			CloudName: request.Avatar.CloudName,
			Extension: request.Avatar.Extension,
			Height:    int(request.Avatar.Height),
			Width:     int(request.Avatar.Width),
		},
		Settings: &common.Settings{
			Theme:    request.Settings.Theme,
			Language: request.Settings.GetLanguage(),
		},
	}, int(request.GetUserId()))

	fmt.Println("OK 2")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "method UserUpdateProfile has something error %s", err)
	}

	fmt.Println("OK 3")

	userInfo, err := storage.FindById(ctx, int(request.UserId))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "method UserUpdateProfile has something error %s", err)
	}

	fmt.Println("OK 4")

	return &proto.UserUpdateProfileResponse{
		UserId: int32(userInfo.Id),
		Email:  *userInfo.Email,
		Phone:  *userInfo.Phone,
		Avatar: &proto.Image{
			Url:       userInfo.Avatar.Url,
			CloudName: userInfo.Avatar.CloudName,
			Extension: userInfo.Avatar.Extension,
			Height:    int32(userInfo.Avatar.Height),
			Width:     int32(userInfo.Avatar.Width),
		},
		Settings: &proto.Settings{
			Theme:    userInfo.Settings.Theme,
			Language: userInfo.Settings.Language,
		},
	}, nil
}

func (r *gRPCServer) UserUpload(ctx context.Context, request *proto.UserUploadRequest) (*proto.UserUploadResponse, error) {
	return nil, nil
}

func (r *gRPCServer) UserUpdatePassword(ctx context.Context, request *proto.UserUpdatePasswordRequest) (*proto.UserUpdatePasswordResponse, error) {
	storage := userstorage.NewSQLStore(r.db)

	err := storage.ChangePassword(ctx, request.GetPassword(), int(request.GetUserId()))

	if err != nil {
		return nil, status.Errorf(codes.Internal, "method UserUpdateProfile has something error %s", err)
	}

	return &proto.UserUpdatePasswordResponse{
		UserId: request.UserId,
		IsDone: true,
	}, nil
}

func (r *gRPCServer) UserUpdateDeviceToken(ctx context.Context, request *proto.UserUpdateDeviceTokenRequest) (*proto.UserUpdateDeviceTokenResponse, error) {
	return nil, nil
}
