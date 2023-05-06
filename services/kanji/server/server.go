// Package server implements all of the gRPC related server request handling.
package server

import (
	"context"
	"fmt"
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
	logger.Info(fmt.Sprintf("fetching kanji (%s)", req.Kanji))
	kanji, err := server.Database.GetKanjiByCharacter(req.Kanji)
	if err != nil {
		return nil, err
	}

	if kanji.WaniKanilevel > 3 {
		// Check auth token validity
		claims, err := tokenValidator.GetClaims(req.AuthToken)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token provided")
		}

		if limit, ok := claims["wanikani_level_limit"]; !ok || int(limit.(float64)) < kanji.WaniKanilevel {
			return nil, status.Error(codes.PermissionDenied, "requires a wanikani subscription for levels higher than 3")
		}
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
	logger.Info(fmt.Sprintf("fetching kanji range (%d-%d)", req.LowerBound, req.UpperBound))
	// Check auth token validity
	if req.UpperBound > 3 {
		claims, err := tokenValidator.GetClaims(req.AuthToken)
		if err != nil {
			return status.Error(codes.Unauthenticated, "valid token required for levels higher than 3")
		}

		if limit, ok := claims["wanikani_level_limit"]; !ok || int32(limit.(float64)) < req.UpperBound {
			return status.Error(codes.PermissionDenied, "requires a wanikani subscription for levels higher than 3")
		}
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
	logger.Info("loading kanji database")
	// Check auth token validity
	jwtClaims, err := tokenValidator.GetClaims(req.AuthToken)
	if err != nil {
		errMsg := "invalid token provided"
		logger.Error(errMsg, err)
		return nil, status.Error(codes.Unauthenticated, errMsg)
	}

	if userType, ok := jwtClaims["type"]; !ok || int(userType.(float64)) != 1 {
		errMsg := "endpoint only valid for admins"
		logger.Info(errMsg)
		return nil, status.Error(codes.PermissionDenied, errMsg)
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
