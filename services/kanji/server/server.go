// Package server implements all of the gRPC related server request handling.
package server

import (
	"context"
	"io"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/common/models"
	tokenValidator "github.com/Aeriqu/kanikaki/common/token"
	"github.com/Aeriqu/kanikaki/services/kanji/database"
	kanjipb "github.com/Aeriqu/kanikaki/services/kanji/proto"
	wanikanipb "github.com/Aeriqu/kanikaki/services/wanikani/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KanjiServer struct {
	kanjipb.UnimplementedKanjiServer
	*database.Database
	wanikanipb.WaniKaniClient
}

func (server *KanjiServer) GetKanji(ctx context.Context, req *kanjipb.KanjiRequest) (*kanjipb.KanjiResponse, error) {
	// Check auth token validity
	_, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token provided")
	}

	kanji, err := server.Database.GetKanjiByCharacter(req.Kanji)
	if err != nil {
		return nil, err
	}

	return &kanjipb.KanjiResponse{
		Character:     kanji.Character,
		WanikaniId:    int32(kanji.WaniKaniId),
		WanikaniLevel: int32(kanji.WaniKanilevel),
		Meanings:      kanji.Meanings,
		Onyomi:        kanji.Onyomi,
		Kunyomi:       kanji.Kunyomi,
		Nanori:        kanji.Nanori,
	}, nil
}

func (server *KanjiServer) GetKanjiByLevelRange(req *kanjipb.KanjiLevelRangeRequest, stream kanjipb.Kanji_GetKanjiByLevelRangeServer) error {
	// Check auth token validity
	_, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		return status.Error(codes.Unauthenticated, "invalid token provided")
	}

	kanjiList, err := server.Database.GetKanjiByRange(int(req.LowerBound), int(req.UpperBound))
	if err != nil {
		return err
	}

	for _, kanji := range *kanjiList {
		response := kanjipb.KanjiResponse{
			Character:     kanji.Character,
			WanikaniId:    int32(kanji.WaniKaniId),
			WanikaniLevel: int32(kanji.WaniKanilevel),
			Meanings:      kanji.Meanings,
			Onyomi:        kanji.Onyomi,
			Kunyomi:       kanji.Kunyomi,
			Nanori:        kanji.Nanori,
		}
		if err := stream.Send(&response); err != nil {
			return err
		}
	}

	return nil
}

func (server *KanjiServer) LoadAllKanji(ctx context.Context, req *kanjipb.WaniKaniTokenRequest) (*kanjipb.LoadKanjiResponse, error) {
	// Check auth token validity
	_, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token provided")
	}

	getAllKanjiRequest := wanikanipb.WaniKaniTokenRequest{
		WanikaniToken: req.WanikaniToken,
		AuthToken:     req.AuthToken,
	}
	getAllKanjiClient, err := server.WaniKaniClient.GetAllKanji(ctx, &getAllKanjiRequest)
	if err != nil {
		errMsg := "error setting up the wanikani client"
		logger.Error(errMsg, err)
		return nil, err
	}

	successList := []string{}
	kanjiList := []models.Kanji{}
	for {
		kanjiData, err := getAllKanjiClient.Recv()
		if err == io.EOF {
			getAllKanjiClient.CloseSend()
			break
		} else if err != nil {
			errMsg := "error obtaining kanji from the wanikani service"
			logger.Error(errMsg, err)
			return nil, status.Error(codes.Aborted, errMsg)
		}

		kanjiList = append(kanjiList, models.Kanji{
			Character:     kanjiData.Character,
			WaniKaniId:    int(kanjiData.WanikaniId),
			WaniKanilevel: int(kanjiData.WanikaniLevel),
			Meanings:      kanjiData.Meanings,
			Onyomi:        kanjiData.Onyomi,
			Kunyomi:       kanjiData.Kunyomi,
			Nanori:        kanjiData.Nanori,
		})

		successList = append(successList, kanjiData.Character)
	}

	_, err = server.Database.AddMultipleKanji(kanjiList)
	if err != nil {
		errMsg := "error adding all kanji to the database"
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Aborted, errMsg)
	}

	return &kanjipb.LoadKanjiResponse{
		CharactersAdded: successList,
	}, nil
}
