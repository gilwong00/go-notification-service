package notificationservice

import (
	db "github.com/gilwong00/go-notification-service/db/sqlc"
	rpcs "github.com/gilwong00/go-notification-service/rpc"
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
