package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/internal/notificationservice"
	"github.com/gilwong00/go-notification-service/internal/userservice"
	config "github.com/gilwong00/go-notification-service/pkg"
	"github.com/gilwong00/go-notification-service/rpcs"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	store := db.NewStore(conn)
	notificationService := notificationservice.NewNotificationService(store)
	userService := userservice.NewUserService(store)
	go startGrpcApiGateway(notificationService, userService, config.HTTPServerAddress)
	startGrpcServices(notificationService, userService, config.GrpcServicesAddress)
}

func startGrpcServices(
	notificationService *notificationservice.NotificationService,
	userService *userservice.UserService,
	address string,
) {
	server := grpc.NewServer()
	rpcs.RegisterNotificationServiceServer(server, notificationService)
	rpcs.RegisterUserServiceServer(server, userService)
	reflection.Register(server)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}
	log.Printf("starting notification service at %s:", listener.Addr().String())
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start notification service:", err)
	}
}

func startGrpcApiGateway(
	notificationService *notificationservice.NotificationService,
	userService *userservice.UserService,
	address string,
) {
	grpcMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: false,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := rpcs.RegisterNotificationServiceHandlerServer(ctx, grpcMux, notificationService)
	if err != nil {
		log.Fatal("cannot register notification service handler", err)
	}
	err = rpcs.RegisterUserServiceHandlerServer(ctx, grpcMux, userService)
	if err != nil {
		log.Fatal("cannot register user service handler", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	log.Printf("gRPC http gateway server starting at %s:", listener.Addr().String())
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("cannot start gRPC http gateway server", err)
	}
}
