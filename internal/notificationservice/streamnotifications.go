package notificationservice

import (
	"context"
	"fmt"

	rpcs "github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
)

func (s *NotificationService) StreamNotifications(
	ctx context.Context,
	req *rpcs.StreamNotificationsRequest,
) (*rpcs.StreamNotificationsResponse, error) {
	return nil, fmt.Errorf("unimplemented rpc: %v", codes.Unimplemented)
}
