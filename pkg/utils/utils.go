package utils

import(
	"fmt"
	"context"
	"encoding/json"

	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

var(
	STATUS_CODE_400 = "Bad Request"
	STATUS_CODE_403 = "Forbidden"
	STATUS_CODE_500 = "Internal Server Error"
)

func ContextValueChecksum(ctx context.Context, keys ...string) (claims map[string]string, errors []*model.Errors) {
	if len(keys) < 1 {
		return
	}

	var value string
	var ok bool
	claims = make(map[string]string)

	for _, key := range keys {
		if value, ok = ctx.Value(key).(string); !ok {
			errors = append(errors, MakeError(500, fmt.Sprintf("Faild get HTTP header value: %s", key)))
		}
		if value == "" {
			errors = append(errors, MakeError(400, fmt.Sprintf("HTTP header value is empty: %s", key)))
		}
		claims[key] = value
	}
	return
}

func MakeError(code int, msg string) (errors *model.Errors) {
	switch code {
	case 400:
		errors = &model.Errors{
			Code: 400,
			Message: STATUS_CODE_400,
			Description: msg,
		}
	case 403:
		errors = &model.Errors{
			Code: 403,
			Message: STATUS_CODE_403,
			Description: msg,
		}
	case 500:
		errors = &model.Errors{
			Code: 500,
			Message: STATUS_CODE_500,
			Description: msg,
		}
	}
	return
}

func MakeErrors(code int, msg string) ([]*model.Errors) {
	var _error *model.Errors
	switch code {
	case 400:
		_error = &model.Errors{
			Code: 400,
			Message: STATUS_CODE_400,
			Description: msg,
		}
	case 403:
		_error = &model.Errors{
			Code: 403,
			Message: STATUS_CODE_403,
			Description: msg,
		}
	case 500:
		_error = &model.Errors{
			Code: 500,
			Message: STATUS_CODE_500,
			Description: msg,
		}
	}
	return []*model.Errors{
		_error,
	}
}

func PackUserInfo(name string, year , month, day, gender int, genre ...string) (*model.User) {
	var _genre []*string
	for _, v := range genre {
		_genre = append(_genre, CastStringPointer(v))
	}
	return &model.User{
		Name: name,
		Birthday: &model.Date{
			Year: year,
			Month: month,
			Day: day,
		},
		Gender: gender,
		Genre: _genre,
	}
}

func PackArticleHeaderInfo(id, title, image string, tag ...string) (*model.ArticleHeader) {
	var tags []*string
	for _, v := range tag {
		tags = append(tags, CastStringPointer(v))
	}
	return &model.ArticleHeader{
		ID: id,
		Title: title,
		ImagePath: image,
		Tags: tags,
	}
}

func PackLogInfo(id, title, date string, time int, concent []float64) (*model.LogInfo) {
	return &model.LogInfo{
		ID: id,
		Title: title,
		Date: date,
		Worktime: time,
		Concentration: concent,
	}
}

func PackLogData(id, title string) (*model.LogData) {
	return &model.LogData{
		ID: id,
		Title: title,
	}
}

func PackSpotInfo(id, name string, latitude, longitude float64) (*model.Spot) {
	return &model.Spot{
		ID: id,
		Name: name,
		Locate: &model.Locate{
			Latitude: latitude,
    		Longitude: longitude,
		},
	}
}

func PackDetourInfo(id, name, image, description string, latitude, longitude float64) (*model.Detour) {
	return &model.Detour{
		ID: id,
		Name: name,
		Image: image,
		Description: description,
		Locate: &model.Locate{
			Latitude: latitude,
			Longitude: longitude,
		},
	}
}

func MakeArticlesRequestJSON(genre string, year, month *int) ([]byte, error) {
	request := model.ArticlesRequest {
		Genre: genre,
		Year: year,
		Month: month,
	}
	return json.Marshal(request)
}

func MakeArticlesResponseStruct(response []byte) (info model.ArticlesResponse) {
	json.Unmarshal(response, &info)
	return
}

func MakeArticlesFromTagResponseJSON(tag string) ([]byte, error) {
	request := model.ArticlesFromTagResponse {
		Tag: tag,
	}
	return json.Marshal(request)
}

func MakeArticleRequestJSON(id string) ([]byte, error) {
	request := model.ArticleRequest {
		ID: id,
	}
	return json.Marshal(request)
}

func MakeArticleResponseStruct(response []byte) (info model.ArticleInfo) {
	json.Unmarshal(response, &info)
	return	
}

func MakeAddLikeResponseStruct(response []byte) (info model.AddLikeResponse) {
	json.Unmarshal(response, &info)
	return
}

func MakeSpotRequestJSON(latitude, longitude float64, worktime, emotion int) ([]byte, error) {
	request := model.SpotRequest {
		Latitude: latitude,
		Longitude: longitude,
		Walktime: worktime,
		Emotion: emotion,
	}
	return json.Marshal(request)
}

func MakeSpotResponseStruct(response []byte) (info model.SpotResponse) {
	json.Unmarshal(response, &info)
	return
}

func MakeDetourRequestJSON(spotLatitude, spotLongitude, userLatitude, userLongitude float64, worktime, emotion int) ([]byte, error) {
	request := model.DetourRequest {
		SpotLatitude: spotLatitude,
		SpotLongitude: spotLongitude,
		UserLatitude: userLatitude,
		UserLongitude: userLongitude,
		Walktime: worktime,
		Emotion: emotion,
	}
	return json.Marshal(request)
}

func MakeDetourResponseStruct(response []byte) (info model.DetourResponse) {
	json.Unmarshal(response, &info)
	return
}

func MakeAddSpotRequest(name, image, desc string, latitude, longitude float64) ([]byte, error) {
	request := model.AddSpotRequest {
		Name: name,
		Image: image,
		Description: desc,
		Latitude: latitude,
		Longitude: longitude,
	}
	return json.Marshal(request)
}

func MakeMutationResponseStruct(response []byte) (info model.AddSpotResponse) {
	json.Unmarshal(response, &info)
	return
}

func MakeAddEvaluationRequestJSON(id string, emotion int, status bool) ([]byte, error) {
	request := model.AddEvaluationRequest {
		ID: id,
		Emotion: emotion,
		Status: status,
	}
	return json.Marshal(request)
}

func CastStringPointer(str string) *string {
	return &str
}
