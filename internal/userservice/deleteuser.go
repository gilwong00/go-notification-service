package userservice

import (
	"context"
	"fmt"

	"github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
)

func (s *UserService) DeleteUser(
	ctx context.Context,
	req *rpcs.DeleteUserRequest,
) (*rpcs.DeleteUserResponse, error) {
	return nil, fmt.Errorf("unimplemented rpc: %v", codes.Unimplemented)
}
