package notificationservice

import (
	"context"
	"fmt"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/pkg/converter"
	rpcs "github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *NotificationService) EnqueueNotification(
	ctx context.Context,
	req *rpcs.EnqueueNotificationRequest,
) (*rpcs.EnqueueNotificationResponse, error) {
	// validation
	if len(req.UserId) == 0 {
		return nil, status.Error(codes.InvalidArgument, "user id is required: %v")
	}
	if len(req.Message) == 0 {
		return nil, status.Error(codes.InvalidArgument, "message is required: %v")
	}
	// get followers of user
	userID, err := converter.StringToUUID(req.UserId)
	if err != nil {
		return nil, err
	}
	followers, err := s.getUserFollowers(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving followers: %v", err)
	}
	for _, follower := range followers {
		if err := s.store.ExecTx(ctx, func(tx *db.Queries) error {
			var err error
			// create initial notification state
			state, err := tx.CreateNotificationState(ctx, db.StatePending)
			if err != nil {
				fmt.Printf("failed to create notification state for follower: %v, err: %v", follower.ID, err)
				return err
			}
			// enqueue notification event
			_, err = tx.CreateNotificationEvent(ctx, db.CreateNotificationEventParams{
				Message:    req.Message,
				FollowerID: follower.ID,
				StateID:    state.ID,
			})
			return err
		}); err != nil {
			fmt.Printf("failed to create notification event for follower: %v, err: %v", follower.ID, err)
			return nil, err
		}
	}
	return &rpcs.EnqueueNotificationResponse{}, nil
}
