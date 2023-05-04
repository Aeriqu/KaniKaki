package main

import (
	"fmt"
	"net"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/kanji/database"
	kanjipb "github.com/Aeriqu/kanikaki/services/kanji/proto"
	"github.com/Aeriqu/kanikaki/services/kanji/server"
	wanikanipb "github.com/Aeriqu/kanikaki/services/wanikani/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	port := ":8080"

	logger.Info(fmt.Sprintf("starting listener on port %v", port))
	portListener, portError := net.Listen("tcp4", port)
	if portError != nil {
		logger.Fatal(fmt.Sprintf("error listening on port %v", port), portError)
	}

	logger.Info("setting up connection with the wanikani service")
	wanikaniLocation := "wanikani.kanikaki.svc.cluster.local:80"
	// TODO: Figure out transport security
	wanikaniConnection, err := grpc.Dial(
		wanikaniLocation,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		errMsg := fmt.Sprintf("error creating connection to: %q", wanikaniLocation)
		logger.Fatal(errMsg, err)
	}
	wanikaniClient := wanikanipb.NewWaniKaniClient(wanikaniConnection)

	grpcServer := grpc.NewServer()
	db := database.Init(&database.ProviderMongodb{})
	kanjipb.RegisterKanjiServer(
		grpcServer,
		&server.KanjiServer{
			UnimplementedKanjiServer: kanjipb.UnimplementedKanjiServer{},
			Database:                 db,
			WaniKaniClient:           wanikaniClient,
		},
	)
	logger.Info(fmt.Sprintf("starting kanji grpc server on %v", portListener.Addr()))

	if serveError := grpcServer.Serve(portListener); serveError != nil {
		logger.Fatal("failed to start grpc server", serveError)
	}
}
