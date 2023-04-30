package clients

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
	helloworldpb "github.com/Aeriqu/kanikaki/services/hello-world/proto"
)

var helloWorldClientInstance *HelloWorldClient
var helloWorldOnce sync.Once

type HelloWorldClient struct {
	grpcClient helloworldpb.HelloWorldClient
}

func (client *HelloWorldClient) SendHelloRequest(name string) (model.HelloWorldQueryResponse, error) {
	request := &helloworldpb.HelloRequest{
		Name: name,
	}
	requestContext, requestCancel := context.WithTimeout(context.Background(), time.Second)
	defer requestCancel()
	helloResponse, helloError := client.grpcClient.Hello(requestContext, request)
	if helloError != nil {
		log.Printf("[HelloWorldClient - SendHelloRequest] Error sending request with name: %q", name)
		log.Println(helloError)
		return model.HelloWorldQueryResponse{}, helloError
	}

	return model.HelloWorldQueryResponse{
		Message: helloResponse.Message,
	}, nil
}

func (client *HelloWorldClient) SendChangeHelloRequest(greeting string) (model.ChangeHelloWorldMutationResponse, error) {
	request := &helloworldpb.ChangeHelloRequest{
		Greeting: greeting,
	}
	requestContext, requestCancel := context.WithTimeout(context.Background(), time.Second)
	defer requestCancel()
	helloResponse, helloError := client.grpcClient.ChangeHello(requestContext, request)
	if helloError != nil {
		log.Printf("[HelloWorldClient - SendChangeHelloRequest] Error sending request with greeting: %q", greeting)
		log.Println(helloError)
		return model.ChangeHelloWorldMutationResponse{}, helloError
	}

	return model.ChangeHelloWorldMutationResponse{
		GreetingChanged: helloResponse.GreetingChanged,
	}, nil
}

func GetHelloWorldClient() *HelloWorldClient {
	helloWorldOnce.Do(func() {
		helloWorldClientInstance = &HelloWorldClient{
			grpcClient: helloworldpb.NewHelloWorldClient(
				getConnection("hello-world-service.kanikaki.svc.cluster.local:80"),
			),
		}
	})
	return helloWorldClientInstance
}
