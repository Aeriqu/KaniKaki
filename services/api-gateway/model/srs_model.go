package model

type IdentifierRequest struct {
	Identifier *string
}

type ReviewDataRequest struct {
	Identifier *string
	Subject    *string
	Topic      *string
	Level      *int32
	ReviewTime *float64
}

type KanjiReviewDataRequest struct {
	Identifier *string
	Character  *string
	SrsLevel   *int32
	ReviewTime *float64
}

type ReviewDataResponse struct {
	Topic      string
	Subject    string
	Level      int32
	ReviewTime float64
}
