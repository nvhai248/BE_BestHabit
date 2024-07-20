package usergrpcclient

import (
	"bestHabit/common"
	proto "bestHabit/generatedProto/proto/userservice"
	"bestHabit/modules/user/usermodel"
	"context"
)

type gRPCUserUpdateProfileClient struct {
	client proto.UserServiceClient
}

func NewGRPCUserUpdateProfileClient(client proto.UserServiceClient) *gRPCUserUpdateProfileClient {
	return &gRPCUserUpdateProfileClient{client: client}
}

func (c *gRPCUserUpdateProfileClient) UpdateUserInfoByGRPC(ctx context.Context, userId int, userUpdate *usermodel.UserUpdate) (*usermodel.User, error) {
	res, err := c.client.UserUpdateProfile(ctx, &proto.UserUpdateProfileRequest{
		UserId: int32(userId),
		Phone:  *userUpdate.Phone,
		Name:   *userUpdate.Name,
		Avatar: &proto.Image{
			Id:        int32(userUpdate.Avatar.Id),
			Url:       userUpdate.Avatar.Url,
			CloudName: userUpdate.Avatar.CloudName,
			Extension: userUpdate.Avatar.Extension,
			Width:     int32(userUpdate.Avatar.Width),
			Height:    int32(userUpdate.Avatar.Height),
		},
		Settings: &proto.Settings{
			Theme:    userUpdate.Settings.Theme,
			Language: userUpdate.Settings.Language,
		},
	})

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return &usermodel.User{
		SQLModel: common.SQLModel{
			Id: int(res.UserId),
		},
		Name:  &res.Name,
		Email: &res.Email,
		FbID:  &res.FbId,
		GgID:  &res.GgId,
		Avatar: &common.Image{
			Url:       res.Avatar.Url,
			CloudName: res.Avatar.CloudName,
			Extension: res.Avatar.Extension,
			Height:    int(res.Avatar.Height),
			Width:     int(res.Avatar.Width),
		},
		Settings: &common.Settings{
			Theme:    res.Settings.Theme,
			Language: res.Settings.GetLanguage(),
		},
	}, nil
}
