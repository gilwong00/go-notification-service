package notificationservice

import (
	"context"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/rpcs"
	"github.com/gofrs/uuid/v5"
)

type NotificationService struct {
	rpcs.UnimplementedNotificationServiceServer
	store db.Store
}

func NewNotificationService(store db.Store) *NotificationService {
	return &NotificationService{
		store: store,
	}
}

func (s *NotificationService) getUserFollowers(
	ctx context.Context,
	userID uuid.UUID,
) ([]db.GetFollowersByUserIDRow, error) {
	var pgFollowers []db.GetFollowersByUserIDRow
	if err := s.store.ExecTx(ctx, func(tx *db.Queries) error {
		var err error
		pgFollowers, err = tx.GetFollowersByUserID(ctx, userID)
		return err
	}); err != nil {
		return nil, err
	}
	return pgFollowers, nil
}
