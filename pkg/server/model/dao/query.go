package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* Access Record API */
// articles
type articles struct {
	Genre string `json:"genre"`
	Year  *int   `json:"Year"`
	Month *int   `json:"month"`
}

func MakeArticlesClient(genre string, year, month *int) AccessResourceRepository {
	return &articles{Genre: genre, Year: year, Month: month}
}
func (info *articles) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf(RECORD_API_QUERY_ARTICLES, pathOf.Record),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

// articlesFromTag
type articlesFromTag struct {
	Tag string `json:"tag"`
}

func MakeArticlesFromTagClient(tag string) AccessResourceRepository {
	return &articlesFromTag{Tag: tag}
}
func (info *articlesFromTag) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf(RECORD_API_QUERY_TAG_ARTICLES, pathOf.Record),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

// article
type article struct {
	ID string `json:"articleID"`
}

func MakeArticleClient(id string) AccessResourceRepository {
	return &article{ID: id}
}
func (info *article) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf(RECORD_API_QUERY_ARTICLE, pathOf.Record),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("UserID", arg[0])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

/* Access Collection API */
type list struct {
	//UserID string `json:"userID"`
}

func MakeListClient() AccessResourceRepository {
	return &list{}
}
func (info *list) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf(COLLECTION_API_QUERY_READLIST, pathOf.Collection),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("UserID", arg[0])
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

/* Access Log API */
type logRequest struct {
}

func NewLogClient() AccessResourceRepository {
	return &logRequest{}
}

func (info *logRequest) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf(LOG_API_QUERY_LOG, pathOf.Log),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("x-token", arg[0])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

type logs struct {
}

func NewLogsClient() AccessResourceRepository {
	return &logs{}
}

func (info *logs) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"GET",
		fmt.Sprintf(LOG_API_QUERY_LOGS, pathOf.Log),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("x-token", arg[0])

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

/* Access Spot API */
// spot
type spot struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Walktime  int     `json:"walktime"`
	Emotion   int     `json:"emotion"`
}

func MakeSpotClient(latitude, longitude float64, times, emotion int) AccessResourceRepository {
	return &spot{Latitude: latitude, Longitude: longitude, Walktime: times, Emotion: emotion}
}
func (info *spot) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(SPOT_API_QUERY_SPOT, pathOf.Spot),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}

// detour
type detour struct {
	SpotLatitude  float64 `json:"spot_latitude"`
	SpotLongitude float64 `json:"spot_longitude"`
	UserLatitude  float64 `json:"user_latitude"`
	UserLongitude float64 `json:"user_longitude"`
	Walktime      int     `json:"walktime"`
	Emotion       int     `json:"emotion"`
}

func MakeDetourClient(sLatitude, sLongitude, uLatitude, uLongitude float64, times, emotion int) AccessResourceRepository {
	return &detour{
		SpotLatitude:  sLatitude,
		SpotLongitude: sLongitude,
		UserLatitude:  uLatitude,
		UserLongitude: uLongitude,
		Walktime:      times,
		Emotion:       emotion,
	}
}
func (info *detour) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(SPOT_API_QUERY_DETOUR, pathOf.Spot),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(response.Body)
}
