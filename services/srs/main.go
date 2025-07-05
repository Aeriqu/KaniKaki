package main

import (
	"fmt"
	"net"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/srs/database"
	srspb "github.com/Aeriqu/kanikaki/services/srs/proto"
	"github.com/Aeriqu/kanikaki/services/srs/server"
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
	srspb.RegisterSrsServer(
		grpcServer,
		&server.SrsServer{
			UnimplementedSrsServer: srspb.UnimplementedSrsServer{},
			Database: db,
		},
	)
	logger.Info(fmt.Sprintf("starting wanikani grpc server on %v", portListener.Addr()))

	if serveError := grpcServer.Serve(portListener); serveError != nil {
		logger.Fatal("failed to start grpc server", serveError)
	}
}
