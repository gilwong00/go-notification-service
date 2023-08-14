package userservice

import (
	"context"
	"fmt"

	"github.com/gilwong00/go-notification-service/rpcs"
	"google.golang.org/grpc/codes"
)

func (s *UserService) CreateUser(
	ctx context.Context,
	req *rpcs.CreateUserRequest,
) (*rpcs.CreateUserResponse, error) {
	return nil, fmt.Errorf("unimplemented rpc: %v", codes.Unimplemented)
}
