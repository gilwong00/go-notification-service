package userservice

import (
	"context"
	"time"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	req *rpcs.CreateUserRequest,
) (*rpcs.CreateUserResponse, error) {
	// validate
	if len(req.Username) == 0 {
		return nil, status.Error(codes.InvalidArgument, "username is required: %v")
	}
	var pgUser db.User
	if err := s.store.ExecTx(ctx, func(q *db.Queries) error {
		var err error
		pgUser, err = q.CreateUser(ctx, req.GetUsername())
		return err
	}); err != nil {
		return nil, err
	}
	createdAt, err := timeToProtoTime(pgUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &rpcs.CreateUserResponse{
		User: &rpcs.User{
			Id:        pgUser.ID.String(),
			Username:  pgUser.Username,
			CreatedAt: createdAt,
		},
	}, nil
}

func timeToProtoTime(t time.Time) (*timestamppb.Timestamp, error) {
	prototime := timestamppb.New(t)
	if err := prototime.CheckValid(); err != nil {
		return nil, err
	}
	return prototime, nil
}
