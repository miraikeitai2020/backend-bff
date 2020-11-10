package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* Access Record API */
// addLike
type addLike struct {
	ID string `json:"articleID"`
}

func MakeAddLikeClient(id string) AccessResourceRepository {
	return &addLike{ID: id}
}
func (info *addLike) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(RECORD_API_MUTATION_LIKE, pathOf.Record),
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

/* Access Memory API */
type addComment struct {
	ID      string `json:"articleID"`
	Comment string `json:"contents"`
}

func MakeAddCommentClient(id, comment string) AccessResourceRepository {
	return &addComment{ID: id, Comment: comment}
}
func (info *addComment) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(MEMORY_API_MUTATION_COMMENT, pathOf.Memory),
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

type addRequest struct {
	Genre    string `json:"genre"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

func MakeAddRequestClient(genre string, year, month int, title, contents string) AccessResourceRepository {
	return &addRequest{Genre: genre, Year: year, Month: month, Title: title, Contents: contents}
}

func (info *addRequest) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(MEMORY_API_MUTATION_REQUEST, pathOf.Memory),
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

/* Access Log API */
type addLog struct {
	Name    string    `json:"logname"`
	Date    string    `json:"datetime"`
	Wrok    int       `json:"worktime"`
	Concent []float64 `json:"concentration"`
}

func NewAddLogClient(name, date string, work int, concent []float64) AccessResourceRepository {
	return &addLog{Name: name, Date: date, Wrok: work, Concent: concent}
}

func (info *addLog) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
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

/* Access Spot API */
type addSpot struct {
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

func MakeAddSpotClient(name, image, desc string, latitude, longitude float64) AccessResourceRepository {
	return &addSpot{Name: name, Image: image, Description: desc, Latitude: latitude, Longitude: longitude}
}
func (info *addSpot) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(SPOT_API_MUTATION_ADD_SPOT, pathOf.Spot),
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

/* Access Evaluation API */
type evaluation struct {
	ID      string `json:"id"`
	Emotion int    `json:"emotion"`
	Status  bool   `json:"status"`
}

func MakeEvaluationClient(id string, emotion int, stat bool) AccessResourceRepository {
	return &evaluation{ID: id, Emotion: emotion, Status: stat}
}
func (info *evaluation) Request(arg ...string) ([]byte, error) {
	body, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(
		"POST",
		fmt.Sprintf(SPOT_API_MUTATION_EVALUATION, pathOf.Evaluation),
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
