// Package server implements all of the gRPC related server request handling.
package server

import (
	tokenValidator "github.com/Aeriqu/kanikaki/common/token"
	"github.com/Aeriqu/kanikaki/services/wanikani/api"
	wanikanipb "github.com/Aeriqu/kanikaki/services/wanikani/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WaniKaniServer struct {
	wanikanipb.UnimplementedWaniKaniServer
}

func (s *WaniKaniServer) GetAllKanji(req *wanikanipb.WaniKaniTokenRequest, stream wanikanipb.WaniKani_GetAllKanjiServer) error {
	// Check auth token validity
	_, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		return status.Error(codes.Unauthenticated, "invalid token provided")
	}
	
	kanjiList, err := api.GetAllKanji(req.WanikaniToken)
	if err != nil {
		return err
	}

	for _, kanji := range kanjiList {
		response := wanikanipb.KanjiResponse{
			Character: kanji.Character,
			WanikaniId: int32(kanji.WaniKaniId),
			WanikaniLevel: int32(kanji.WaniKanilevel),
			Meanings: kanji.Meanings,
			Onyomi: kanji.Onyomi,
			Kunyomi: kanji.Kunyomi,
			Nanori: kanji.Nanori,
		}
		if err := stream.Send(&response); err !=  nil {
			return err
		}
	}

	return nil
}
