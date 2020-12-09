package dto

import (
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

type ArticlesResponse struct {
	ArticleList []*model.ArticleHeader `json:"articles"`
}

type ArticleResponse struct {
	ArticleInfo model.ArticleInfo `json:"articleDetail"`
}


type ListResponse struct {
	Articles []*model.ArticleHeader `json:"Lists"`
}

type LogResponse struct {
	Info model.LogInfo
}

type LogsResponse struct {
	Data []*model.LogData
}

type SpotResponse struct {
	Spot SpotInfo `json:"spot"`
}

type SpotInfo struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type DetourResponse struct {
	Detour []DetourInfo `json:"detours"`
}

type DetourInfo struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}
