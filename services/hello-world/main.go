package main

import (
	"log"
	"net"

	helloworldpb "github.com/Aeriqu/kanikaki/services/hello-world/proto"
	"github.com/Aeriqu/kanikaki/services/hello-world/server"
	"google.golang.org/grpc"
)

func main() {
	port := ":8080"

	log.Printf("starting listener on port %v", port)
	portListener, portError := net.Listen("tcp4", port)
	if portError != nil {
		log.Fatalf("error trying to listen on port %v %v", port, portError)
	}

	grpcServer := grpc.NewServer()
	helloworldpb.RegisterHelloWorldServer(grpcServer, &server.HelloWorldServer{})
	log.Printf("started grpc server on %v", portListener.Addr())

	if serveError := grpcServer.Serve(portListener); serveError != nil {
		log.Fatalf("failed to serve: %v", serveError)
	}
}
