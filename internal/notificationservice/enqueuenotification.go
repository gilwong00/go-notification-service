package notificationservice

import (
	"context"
	"fmt"
	"log"

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
	if err := s.store.ExecTx(ctx, func(tx *db.Queries) error {
		// TODO clean this up
		var err error
		for _, follower := range followers {
			// create initial notification state
			state, err := tx.CreateNotificationState(ctx, db.StatePending)
			if err != nil {
				log.Printf("failed to create notification state for follower: %v, err: %v", follower.ID, err)
				return err
			}
			// enqueue notification event if state was created successfully
			_, err = tx.CreateNotificationEvent(ctx, db.CreateNotificationEventParams{
				Message:    req.Message,
				FollowerID: follower.ID,
				StateID:    state.ID,
			})
			if err != nil {
				log.Printf("failed to create notification event for follower: %v, err: %v", follower.ID, err)
				return err
			}
		}
		return err
	}); err != nil {
		return nil, err
	}
	return &rpcs.EnqueueNotificationResponse{}, nil
}
