package model

type KanjiRequest struct {
	Kanji *string
}

type KanjiLevelRangeRequest struct {
	LowerBound *int32
	UpperBound *int32
}

type WaniKaniTokenRequest struct {
	WanikaniToken *string
}

type KanjiResponse struct {
	Character string
	WanikaniId int32
	WanikaniLevel int32
	Meanings *[]string
	Onyomi *[]string
	Kunyomi *[]string
	Nanori *[]string
}

type LoadKanjiResponse struct {
	CharactersAdded *[]string
}