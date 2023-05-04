package api

import "time"

type WaniKaniSubjectResponse struct {
	Object        string          `json:"object"`
	URL           string          `json:"url"`
	Pages         Pages           `json:"pages"`
	TotalCount    int             `json:"total_count"`
	DataUpdatedAt time.Time       `json:"data_updated_at"`
	Data          []KanjiMetadata `json:"data"`
}

type Pages struct {
	PerPage     int    `json:"per_page"`
	NextURL     string `json:"next_url"`
	PreviousURL string `json:"previous_url"`
}

type KanjiMetadata struct {
	ID            int       `json:"id"`
	Object        string    `json:"object"`
	URL           string    `json:"url"`
	DataUpdatedAt time.Time `json:"data_updated_at"`
	Data          KanjiData `json:"data"`
}

type KanjiData struct {
	CreatedAt                 time.Time           `json:"created_at"`
	Level                     int                 `json:"level"`
	Slug                      string              `json:"slug"`
	HiddenAt                  time.Time           `json:"hidden_at"`
	DocumentURL               string              `json:"document_url"`
	Characters                string              `json:"characters"`
	Meanings                  []Meanings          `json:"meanings"`
	AuxiliaryMeanings         []AuxiliaryMeanings `json:"auxiliary_meanings"`
	Readings                  []Readings          `json:"readings"`
	ComponentSubjectIds       []int               `json:"component_subject_ids"`
	AmalgamationSubjectIds    []int               `json:"amalgamation_subject_ids"`
	VisuallySimilarSubjectIds []int               `json:"visually_similar_subject_ids"`
	MeaningMnemonic           string              `json:"meaning_mnemonic"`
	MeaningHint               string              `json:"meaning_hint"`
	ReadingMnemonic           string              `json:"reading_mnemonic"`
	ReadingHint               string              `json:"reading_hint"`
	LessonPosition            int                 `json:"lesson_position"`
	SpacedRepetitionSystemID  int                 `json:"spaced_repetition_system_id"`
}

type Meanings struct {
	Meaning        string `json:"meaning"`
	Primary        bool   `json:"primary"`
	AcceptedAnswer bool   `json:"accepted_answer"`
}

type AuxiliaryMeanings struct {
	Type    string `json:"type"`
	Meaning string `json:"meaning"`
}

type Readings struct {
	Type           string `json:"type"`
	Primary        bool   `json:"primary"`
	Reading        string `json:"reading"`
	AcceptedAnswer bool   `json:"accepted_answer"`
}
