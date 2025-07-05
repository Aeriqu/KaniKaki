package resolver

import (
	"context"
	"errors"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/api-gateway/clients"
	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
)

func (*QueryResolver) GetAllKanjiReviewData(ctx context.Context, query model.IdentifierRequest) ([]model.ReviewDataResponse, error) {
	kanjiReviewResponse, err := clients.GetSrsClient().GetAllKanjiReviewData(ctx, *query.Identifier)
	if err != nil {
		return nil, err
	}

	return kanjiReviewResponse, nil
}

func (*QueryResolver) GetDueKanjiReviews(ctx context.Context, query model.IdentifierRequest) ([]model.ReviewDataResponse, error) {
	kanjiReviewResponse, err := clients.GetSrsClient().GetDueKanjiReviews(ctx, *query.Identifier)
	if err != nil {
		return nil, err
	}

	return kanjiReviewResponse, nil
}

func (*MutationResolver) UpdateReviewData(ctx context.Context, mutation model.ReviewDataRequest) (model.ReviewDataResponse, error) {
	if *mutation.Subject == "kanji" {
		kanjiMutation := model.KanjiReviewDataRequest{
			Identifier: mutation.Identifier,
			Character:  mutation.Topic,
			SrsLevel:   mutation.Level,
			ReviewTime: mutation.ReviewTime,
		}
		kanjiReviewResponse, err := clients.GetSrsClient().UpdateKanjiReviewData(ctx, &kanjiMutation)
		if err != nil {
			return model.ReviewDataResponse{}, err
		}
		return kanjiReviewResponse, nil
	}

	logger.Info("UpdateReviewData called with unsupported subject type")
	return model.ReviewDataResponse{}, errors.New("unsupported subject type")
}