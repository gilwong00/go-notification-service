package notificationservice

import (
	"context"
	"fmt"

	rpcs "github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
)

func (s *NotificationService) EnqueueNotification(
	ctx context.Context,
	req *rpcs.EnqueueNotificationRequest,
) (*rpcs.EnqueueNotificationResponse, error) {
	return nil, fmt.Errorf("unimplemented rpc: %v", codes.Unimplemented)
}
