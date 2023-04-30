package model

type HelloWorldQuery struct {
	Name *string
}

type HelloWorldQueryResponse struct {
	Message string
}

type ChangeHelloWorldMutation struct {
	Greeting *string
}

type ChangeHelloWorldMutationResponse struct {
	GreetingChanged string
}