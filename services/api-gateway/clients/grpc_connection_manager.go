package clients

import (
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcConnectionManagerInstance *GrpcConnectionManager
var grpcOnce sync.Once

type GrpcConnectionManager struct {
	Connections map[string]*grpc.ClientConn
}

func getGrpcConnectionHandler() *GrpcConnectionManager {
	grpcOnce.Do(func() {
		grpcConnectionManagerInstance = &GrpcConnectionManager{
			make(map[string]*grpc.ClientConn),
		}
	})
	return grpcConnectionManagerInstance
}

func getConnection(location string) *grpc.ClientConn {
	var connection *grpc.ClientConn
	manager := getGrpcConnectionHandler()

	if connectionExisting, connectionExists := manager.Connections[location]; connectionExists {
		connection = connectionExisting
	} else {
		connectionCreated, connectionError := grpc.Dial(location,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if connectionError != nil {
			log.Printf("[GrpcClient - getConnection] Error creating connection to: %q", location)
			log.Println(connectionError)
			return nil
		}
		manager.Connections[location] = connectionCreated
		connection = connectionCreated
	}

	return connection
}

func CloseConnections() {
	instance := getGrpcConnectionHandler()
	for _, connection := range instance.Connections {
		connection.Close()
	}
}
