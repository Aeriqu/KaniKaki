package clients

import (
	"fmt"
	"sync"

	"github.com/Aeriqu/kanikaki/common/logger"
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
		// TODO: Figure out transport security
		connectionCreated, err := grpc.NewClient(location,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			errMsg := fmt.Sprintf("error creating connection to: %q", location)
			logger.Error(errMsg, err)
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
