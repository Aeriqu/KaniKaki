package clients

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
	srspb "github.com/Aeriqu/kanikaki/services/srs/proto"
)

var srsClientInstance *SrsClient
var srsClientOnce sync.Once

type SrsClient struct {
	grpcClient srspb.SrsClient
}

func (c *SrsClient) GetAllKanjiReviewData(ctx context.Context, identifier string) ([]model.ReviewDataResponse, error) {
	_, token, _ := validateToken(ctx)

	request := &srspb.IdentifierRequest{
		Identifier: identifier,
		AuthToken: token,
	}

	stream, err := c.grpcClient.GetAllKanjiReviewData(ctx, request)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get all kanji review data for identifier: %s", identifier), err)
		return nil, err
	}

	var responses []model.ReviewDataResponse
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			stream.CloseSend()
			break
		}
		if err != nil {
			logger.Error(fmt.Sprintf("error receiving kanji review data for identifier: %s", identifier), err)
			return nil, err
		}

		responses = append(responses, model.ReviewDataResponse{
			Subject:    "kanji",
			Topic:      response.Character,
			Level:      response.SrsLevel,
			ReviewTime: float64(response.ReviewTime),
		})
	}

	return responses, nil
}

func (c *SrsClient) GetDueKanjiReviews(ctx context.Context, identifier string) ([]model.ReviewDataResponse, error) {
	_, token, _ := validateToken(ctx)

	req := &srspb.IdentifierRequest{
		Identifier: identifier,
		AuthToken:  token,
	}

	stream, err := c.grpcClient.GetDueKanjiReviews(ctx, req)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get due kanji reviews for identifier: %s", identifier), err)
		return nil, err
	}

	var responses []model.ReviewDataResponse
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			stream.CloseSend()
			break
		}
		if err != nil {
			logger.Error(fmt.Sprintf("error receiving due kanji reviews for identifier: %s", identifier), err)
			return nil, err
		}

		responses = append(responses, model.ReviewDataResponse{
			Subject:    "kanji",
			Topic:      response.Character,
			Level:      response.SrsLevel,
			ReviewTime: float64(response.ReviewTime),
		})
	}

	return responses, nil
}

func (c *SrsClient) UpdateKanjiReviewData(ctx context.Context, req *model.KanjiReviewDataRequest) (model.ReviewDataResponse, error) {
	_, token, _ := validateToken(ctx)

	grpcReq := &srspb.KanjiReviewDataRequest{
		Identifier: *req.Identifier,
		Character:  *req.Character,
		SrsLevel:   *req.SrsLevel,
		ReviewTime: int64(*req.ReviewTime),
		AuthToken:  token,
	}

	response, err := c.grpcClient.UpdateKanjiReviewData(ctx, grpcReq)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to update kanji review data for user (%s) for character: %s", *req.Identifier, *req.Character), err)
		return model.ReviewDataResponse{}, err
	}

	return model.ReviewDataResponse{
		Subject:    "kanji",
		Topic:      response.Character,
		Level:      response.SrsLevel,
		ReviewTime: float64(response.ReviewTime),
	}, nil
}

func GetSrsClient() *SrsClient {
	srsClientOnce.Do(func() {
		srsClientInstance = &SrsClient{
			grpcClient: srspb.NewSrsClient(
				getConnection("srs-service.kanikaki.svc.cluster.local:80"),
			),
		}
	})
	return srsClientInstance
}
