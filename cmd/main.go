package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	db "github.com/gilwong00/go-notification-service/db/sqlc"
	"github.com/gilwong00/go-notification-service/internal/notificationservice"
	config "github.com/gilwong00/go-notification-service/pkg"
	rpcs "github.com/gilwong00/go-notification-service/rpc"
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
	go startGrpcApiGateway(notificationService, config.HTTPServerAddress)
	startNotificationService(notificationService, config.NotificationServiceAddress)
}

func startNotificationService(
	service *notificationservice.NotificationService,
	address string,
) {
	server := grpc.NewServer()
	rpcs.RegisterNotificationServiceServer(server, service)
	reflection.Register(server)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Cannot create listener:", err)
	}
	log.Printf("starting notification service at %s:", listener.Addr().String())
	err = server.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start grpc server:", err)
	}
}

func startGrpcApiGateway(
	service *notificationservice.NotificationService,
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
	err := rpcs.RegisterNotificationServiceHandlerServer(ctx, grpcMux, service)
	if err != nil {
		log.Fatal("cannot register notification service handler", err)
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
