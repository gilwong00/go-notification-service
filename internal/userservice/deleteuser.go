package userservice

import (
	"context"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/pkg/converter"
	"github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserService) DeleteUser(
	ctx context.Context,
	req *rpcs.DeleteUserRequest,
) (*rpcs.DeleteUserResponse, error) {
	if len(req.Id) == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id is required: %v")
	}
	userID, err := converter.StringToUUID(req.Id)
	if err != nil {
		return nil, err
	}
	if txErr := s.store.ExecTx(ctx, func(tx *db.Queries) error {
		err := tx.DeleteUser(ctx, userID)
		return err
	}); txErr != nil {
		return nil, err
	}
	return &rpcs.DeleteUserResponse{}, nil
}
