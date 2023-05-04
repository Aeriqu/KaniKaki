package resolver

import (
	"context"

	"github.com/Aeriqu/kanikaki/services/api-gateway/clients"
	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
)

func (*QueryResolver) GetKanji(ctx context.Context, query model.KanjiRequest) (*model.KanjiResponse, error) {
	kanjiResponse, err := clients.GetKanjiClient().GetKanji(ctx, *query.Kanji)
	if err != nil {
		return nil, err
	}

	return &kanjiResponse, nil
}

func (*QueryResolver) GetKanjiByLevelRange(ctx context.Context, query model.KanjiLevelRangeRequest) ([]model.KanjiResponse, error) {
	kanjiResponse, err := clients.GetKanjiClient().GetKanjiByLevelRange(ctx, int(*query.LowerBound), int(*query.UpperBound))
	if err != nil {
		return nil, err
	}

	return kanjiResponse, nil
}

func (*MutationResolver) LoadAllKanji(ctx context.Context, mutation model.WaniKaniTokenRequest) (*model.LoadKanjiResponse, error) {
	kanjiResponse, err := clients.GetKanjiClient().LoadAllKanji(ctx, *mutation.WanikaniToken)
	if err != nil {
		return nil, err
	}

	return &kanjiResponse, nil
}