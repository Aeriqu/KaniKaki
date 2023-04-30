package resolver

import (
	"context"

	"github.com/Aeriqu/kanikaki/services/api-gateway/clients"
	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
)

func (*QueryResolver) HelloWorld(ctx context.Context, query model.HelloWorldQuery) (*model.HelloWorldQueryResponse, error) {
	response := &model.HelloWorldQueryResponse{}
	helloResponse, helloResponseError := clients.GetHelloWorldClient().SendHelloRequest(*query.Name)
	if helloResponseError != nil {
		return response, helloResponseError
	}
	response.Message = helloResponse.Message
	return response, nil
}

func (*MutationResolver) ChangeHelloWorld(ctx context.Context, mutation model.ChangeHelloWorldMutation) (*model.ChangeHelloWorldMutationResponse, error) {
	response := &model.ChangeHelloWorldMutationResponse{}
	helloResponse, helloResponseError := clients.GetHelloWorldClient().SendChangeHelloRequest(*mutation.Greeting)
	if helloResponseError != nil {
		return response, helloResponseError
	}
	response.GreetingChanged = helloResponse.GreetingChanged
	return response, nil
}
