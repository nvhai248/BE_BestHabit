package usergrpcclient

import (
	"bestHabit/common"
	proto "bestHabit/generatedProto/proto/userservice"
	"context"
)

type gRPCUserUpdatePwClient struct {
	client proto.UserServiceClient
}

func NewGRPCUserUpdatePwClient(client proto.UserServiceClient) *gRPCUserUpdatePwClient {
	return &gRPCUserUpdatePwClient{client: client}
}

func (c *gRPCUserUpdatePwClient) UpdatePasswordByGRPC(ctx context.Context, userId int, newPassword string) (*int, error) {
	res, err := c.client.UserUpdatePassword(ctx, &proto.UserUpdatePasswordRequest{
		UserId:   int32(userId),
		Password: newPassword,
	})

	if err != nil {
		return nil, common.ErrDB(err)
	}

	var result = int(res.GetUserId())
	return &result, nil
}
