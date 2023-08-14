package notificationservice

import (
	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/rpcs"
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
