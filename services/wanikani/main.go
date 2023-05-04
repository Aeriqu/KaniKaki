package main

import (
	"fmt"
	"net"

	"github.com/Aeriqu/kanikaki/common/logger"
	wanikanipb "github.com/Aeriqu/kanikaki/services/wanikani/proto"
	"github.com/Aeriqu/kanikaki/services/wanikani/server"
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
	wanikanipb.RegisterWaniKaniServer(
		grpcServer,
		&server.WaniKaniServer{
			UnimplementedWaniKaniServer: wanikanipb.UnimplementedWaniKaniServer{},
		},
	)
	logger.Info(fmt.Sprintf("starting wanikani grpc server on %v", portListener.Addr()))

	if serveError := grpcServer.Serve(portListener); serveError != nil {
		logger.Fatal("failed to start grpc server", serveError)
	}
}
