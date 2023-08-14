package userservice

import (
	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/rpcs"
)

type UserService struct {
	rpcs.UnimplementedUserServiceServer
	store db.Store
}

func NewUserService(store db.Store) *UserService {
	return &UserService{
		store: store,
	}
}
