package main

import (
	"fmt"
	"net"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/auth/database"
	authpb "github.com/Aeriqu/kanikaki/services/auth/proto"
	"github.com/Aeriqu/kanikaki/services/auth/server"
	"google.golang.org/grpc"
)

func main() {
	port := ":8080"

	logger.Info(fmt.Sprintf("starting listener on port %v", port))
	portListener, portError := net.Listen("tcp4", port)
	if portError != nil {
		logger.Fatal(fmt.Sprintf("error listening on port %v", port), portError)
	}

	grpcServer := grpc.NewServer()
	db := database.Init(&database.ProviderMongodb{})
	authpb.RegisterAuthServer(
		grpcServer,
		&server.AuthServer{
			UnimplementedAuthServer: authpb.UnimplementedAuthServer{},
			Database:                db,
		},
	)
	logger.Info(fmt.Sprintf("starting auth grpc server on %v", portListener.Addr()))

	if serveError := grpcServer.Serve(portListener); serveError != nil {
		logger.Fatal("failed to start grpc server", serveError)
	}
}
