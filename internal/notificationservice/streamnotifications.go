package notificationservice

import (
	"database/sql"
	"errors"
	"time"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	rpcs "github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxRetryAttempt int32 = 5

func (s *NotificationService) StreamNotifications(
	req *rpcs.StreamNotificationsRequest,
	server rpcs.NotificationService_StreamNotificationsServer,
) error {
	ctx := server.Context()
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for ctx.Err() == nil {
		// get sendable events
		events, err := s.store.ListSendableNotifications(ctx)
		if err != nil {
			return err
		}
		if len(events) == 0 {
			return errors.New("no sendable notifications")
		}
		for _, event := range events {
			if event.Url.Valid {
				payload := &rpcs.StreamNotificationsResponse{
					Event: &rpcs.NotificationEvent{
						Message:     event.Message,
						FollowerUrl: event.Url.String,
					},
				}
				if event.Attempts.Int32 == maxRetryAttempt {
					if err := s.store.DeleteNotificationEvent(ctx, db.DeleteNotificationEventParams{
						ID:      event.NotificationStateID,
						State:   db.StateFailed,
						Message: formatSqlNullString(event.Message),
					}); err != nil {
						return err
					}
					return errors.New("maximum amount of attempts exceed")
				}
				err := server.Send(payload)
				// update attempt count
				if txErr := s.store.UpdateNotificationAttemptCount(ctx, db.UpdateNotificationAttemptCountParams{
					ID:       event.ID,
					Attempts: sql.NullInt32{Int32: event.Attempts.Int32 + 1, Valid: true},
				}); txErr != nil {
					return txErr
				}
				// if we failed to send
				if err != nil {
					if txErr := s.store.UpdateNotificationStateByID(ctx, db.UpdateNotificationStateByIDParams{
						ID:      event.NotificationStateID,
						State:   db.StateFailed,
						Message: formatSqlNullString(event.Message),
					}); txErr != nil {
						return txErr
					}
					return err
				}
				// update state to success
				if txErr := s.store.UpdateNotificationStateByID(ctx, db.UpdateNotificationStateByIDParams{
					ID:          event.NotificationStateID,
					State:       db.StateSuccess,
					Message:     formatSqlNullString(event.Message),
					CompletedAt: sql.NullTime{Time: time.Now(), Valid: true},
				}); txErr != nil {
					return txErr
				}
			}
		}
		// TODO: use a channel for ticker
	}
	return status.Errorf(codes.Canceled, "context canceled")
}

func formatSqlNullString(str string) sql.NullString {
	if len(str) == 0 {
		return sql.NullString{String: "", Valid: false}
	}
	return sql.NullString{String: str, Valid: true}
}
