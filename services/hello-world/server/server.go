package server

import (
	"context"

	helloworldpb "github.com/Aeriqu/kanikaki/services/hello-world/proto"
)

var greeting string = "Hello"

type HelloWorldServer struct {
	helloworldpb.UnimplementedHelloWorldServer
}

func (s *HelloWorldServer) Hello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloResponse, error) {
	return &helloworldpb.HelloResponse{Message: greeting + " " + in.GetName()}, nil
}

func (s *HelloWorldServer) ChangeHello(ctx context.Context, in *helloworldpb.ChangeHelloRequest) (*helloworldpb.ChangeHelloResponse, error) {
	greeting = in.GetGreeting()
	return &helloworldpb.ChangeHelloResponse{GreetingChanged: greeting}, nil
}
