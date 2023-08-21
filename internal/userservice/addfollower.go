package userservice

import (
	"context"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/pkg/converter"
	"github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserService) AddFollower(
	ctx context.Context,
	req *rpcs.AddFollowerRequest,
) (*rpcs.AddFollowerResponse, error) {
	// validation
	if len(req.UserId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id is required: %v")
	}
	if len(req.FollowerId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "follower id is required: %v")
	}
	userId, err := converter.StringToUUID(req.UserId)
	if err != nil {
		return nil, err
	}
	followerId, err := converter.StringToUUID(req.FollowerId)
	if err != nil {
		return nil, err
	}
	if txErr := s.store.ExecTx(ctx, func(tx *db.Queries) error {
		_, err := tx.CreateFollower(ctx, db.CreateFollowerParams{
			UserID:     userId,
			FollowerID: followerId,
		})
		return err
	}); txErr != nil {
		return nil, txErr
	}
	return &rpcs.AddFollowerResponse{}, nil
}
