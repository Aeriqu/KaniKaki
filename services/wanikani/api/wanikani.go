// Package api implements communication with external apis.
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/common/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var WANIKANI_BASE_URL string = "https://api.wanikani.com"
var WANIKANI_REVISION string = "20170710"

func sendGetRequest(uri string, formFields url.Values, token string) (*http.Response, error) {
	httpClient := http.Client{}
	request, err := http.NewRequest(http.MethodGet, uri, strings.NewReader(formFields.Encode()))
		if err != nil {
			errMsg := "error crafting a response for wanikani"
			logger.Error(errMsg, err)
			return nil, err
		}
		request.Header.Add("Authorization", "Bearer "+token)
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		request.Header.Add("Wanikani-Revision", WANIKANI_REVISION)

		response, err := httpClient.Do(request)
		if err != nil {
			errMsg := "error getting a response from wanikani"
			logger.Error(errMsg, err)
			return nil, err
		}

		return response, nil
}

func processSubjectKanjiData(responseData WaniKaniSubjectResponse) (string, []models.Kanji) {
	result := []models.Kanji{}
	for _, item := range responseData.Data {
		newKanji := models.Kanji{
			Character: item.Data.Characters,
			WaniKaniId: item.ID,
			WaniKanilevel: item.Data.Level,
			Meanings: []string{},
			Onyomi: []string{},
			Kunyomi: []string{},
			Nanori: []string{},
		}

		for _, meaning := range item.Data.Meanings {
			newKanji.Meanings = append(newKanji.Meanings, meaning.Meaning)
		}

		for _, reading := range item.Data.Readings {
			if reading.Type == "onyomi" {
				newKanji.Onyomi = append(newKanji.Onyomi, reading.Reading)
			} else if reading.Type == "kunyomi" {
				newKanji.Kunyomi = append(newKanji.Kunyomi, reading.Reading)
			} else if reading.Type == "nanori" {
				newKanji.Nanori = append(newKanji.Nanori, reading.Reading)
			}
		}

		result = append(result, newKanji)
	}
	return responseData.Pages.NextURL, result
}

func GetAllKanji(token string) ([]models.Kanji, error) {
	result := []models.Kanji{}

	logger.Info("attempting to get all kanji")
	if len(token) == 0 {
		return result, status.Error(codes.Aborted, "no token provided")
	}

	nextUri := WANIKANI_BASE_URL + "/v2/subjects"
	formFields :=  url.Values{}
	formFields.Set("types", "kanji")
	

	for nextUri != "" {
		response, err := sendGetRequest(nextUri, formFields, token)
		if err != nil {
			errMsg := "error getting an all kanji response from wanikani"
			logger.Error(errMsg, err)
			return nil, status.Error(codes.Aborted, errMsg)
		}

		var responseData WaniKaniSubjectResponse
		err = json.NewDecoder(response.Body).Decode(&responseData)
		if err != nil {
			errMsg := "error parsing the response from wanikani"
			logger.Error(errMsg, err)
			return result, status.Error(codes.Aborted, errMsg)
		}

		newUri, processedKanji := processSubjectKanjiData(responseData)
		nextUri = newUri
		result = append(result, processedKanji...)
	}

	logger.Info(fmt.Sprintf("found %d kanji", len(result)))

	return result, nil
}