package clients

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
	kanjipb "github.com/Aeriqu/kanikaki/services/kanji/proto"
)

var kanjiClientInstance *KanjiClient
var kanjiClientOnce sync.Once

type KanjiClient struct {
	grpcClient kanjipb.KanjiClient
}

// GetKanji sends a request to the kanji service to grab a given kanji
func (client *KanjiClient) GetKanji(ctx context.Context, kanji string) (model.KanjiResponse, error) {
	_, token, _ := validateToken(ctx)

	request := &kanjipb.KanjiRequest{
		Kanji:     kanji,
		AuthToken: token,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	kanjiResponse, err := client.grpcClient.GetKanji(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending request with kanji (%s)", kanji), err)
		return model.KanjiResponse{}, err
	}

	return model.KanjiResponse{
		Character:     kanjiResponse.Character,
		WanikaniId:    kanjiResponse.WanikaniId,
		WanikaniLevel: kanjiResponse.WanikaniLevel,
		Meanings:      &kanjiResponse.Meanings,
		Onyomi:        &kanjiResponse.Onyomi,
		Kunyomi:       &kanjiResponse.Kunyomi,
		Nanori:        &kanjiResponse.Nanori,
	}, nil
}

// GetKanjiByLevelRange sends a request to the kanji service to grab kanji
// by a given level range.
func (client *KanjiClient) GetKanjiByLevelRange(ctx context.Context, lowerBound int, upperBound int) ([]model.KanjiResponse, error) {
	_, token, _ := validateToken(ctx)

	request := &kanjipb.KanjiLevelRangeRequest{
		LowerBound: int32(lowerBound),
		UpperBound: int32(upperBound),
		AuthToken:  token,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	kanjiClient, err := client.grpcClient.GetKanjiByLevelRange(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending kanji request with bounds (%d-%d)", lowerBound, upperBound), err)
		return []model.KanjiResponse{}, err
	}

	kanjiList := []model.KanjiResponse{}

	for {
		kanjiData, err := kanjiClient.Recv()
		if err == io.EOF {
			kanjiClient.CloseSend()
			break
		} else if err != nil {
			logger.Error("error obtaining kanji from the kanji service", err)
			return nil, err
		}

		kanjiList = append(kanjiList, model.KanjiResponse{
			Character:     kanjiData.Character,
			WanikaniId:    kanjiData.WanikaniId,
			WanikaniLevel: kanjiData.WanikaniLevel,
			Meanings:      &kanjiData.Meanings,
			Onyomi:        &kanjiData.Onyomi,
			Kunyomi:       &kanjiData.Kunyomi,
			Nanori:        &kanjiData.Nanori,
		})
	}

	return kanjiList, nil
}

// LoadAllKanji sends a request to the kanji service to perform its initial load
// of all of the kanji in wanikani
func (client *KanjiClient) LoadAllKanji(ctx context.Context, wanikaniToken string) (model.LoadKanjiResponse, error) {
	_, token, err := validateToken(ctx)
	if err != nil {
		return model.LoadKanjiResponse{}, err
	}

	request := &kanjipb.WaniKaniTokenRequest{
		WanikaniToken: wanikaniToken,
		AuthToken:     token,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second*5)
	defer requestCancel()
	kanjiResponse, err := client.grpcClient.LoadAllKanji(requestContext, request)
	if err != nil {
		logger.Error("error sending request to obtain all kanji", err)
		return model.LoadKanjiResponse{}, err
	}

	return model.LoadKanjiResponse{
		CharactersAdded: &kanjiResponse.CharactersAdded,
	}, nil
}

func GetKanjiClient() *KanjiClient {
	kanjiClientOnce.Do(func() {
		kanjiClientInstance = &KanjiClient{
			grpcClient: kanjipb.NewKanjiClient(
				getConnection("kanji-service.kanikaki.svc.cluster.local:80"),
			),
		}
	})
	return kanjiClientInstance
}
