package service

import (
	"encoding/json"

	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/dto"
)

func ConvertResponseArticles(response []byte) (info dto.ArticlesResponse) {
	json.Unmarshal(response, &info)
	return
}

func ConvertResponseArticle(response []byte) (info dto.ArticleResponse) {
	json.Unmarshal(response, &info.ArticleInfo)
	return
}

func ConvertResponseSpot(response []byte) (info dto.SpotResponse) {
	json.Unmarshal(response, &info)
	return
}

func ConvertResponseList(response []byte) (info dto.ListResponse) {
	var articleList []model.ArticleHeader
	json.Unmarshal(response, &articleList)
	for i := 0; i < len(articleList); i++ {
		info.Articles = append(info.Articles, &articleList[i])
	}
	return
}

func ConvertResponseLog(response []byte) (info dto.LogResponse) {
	var logInfo struct {
		Title         string    `json:"LogName"`
		Date          string    `json:"Date"`
		Worktime      int       `json:"WorkTime"`
		Concentration []float64 `json:"Concentration"`
	}
	json.Unmarshal(response, &logInfo)
	info.Info = model.LogInfo{
		Title:         logInfo.Title,
		Date:          logInfo.Date,
		Worktime:      logInfo.Worktime,
		Concentration: logInfo.Concentration,
	}
	return
}

func ConvertResponseDetour(response []byte) (info dto.DetourResponse) {
	json.Unmarshal(response, &info)
	return
}

func ConvertResponseLike(response []byte) dto.MutationResponse {
	var info dto.LikeResponse
	var status bool
	json.Unmarshal(response, &info)
	if info.Nice == 100 {
		status = false
	} else {
		status = true
	}
	return dto.MutationResponse{Status: status}
}

func ConvertResponseMutation(response []byte) (info dto.MutationResponse) {
	json.Unmarshal(response, &info)
	return
}

// type converter
func ConvertSpotDtoToModel(d dto.SpotResponse) *model.Spot {
	return &model.Spot{
		ID:   d.Spot.ID,
		Name: d.Spot.Name,
		Locate: &model.Locate{
			Latitude:  d.Spot.Latitude,
			Longitude: d.Spot.Longitude,
		},
	}
}

func ConvertDetourDtoToModel(d dto.DetourResponse) (detour []*model.Detour) {
	for _, v := range d.Detour {
		detour = append(
			detour,
			&model.Detour{
				ID:          v.ID,
				Name:        v.Name,
				Image:       v.Image,
				Description: v.Description,
				Locate: &model.Locate{
					Latitude:  v.Latitude,
					Longitude: v.Longitude,
				},
			},
		)
	}
	return
}
