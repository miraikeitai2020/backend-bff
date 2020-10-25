package utils

import(
	"fmt"
	"context"

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

func PackArticleInfo(id, title, image string, niceNum int, ctxt string, nice, list bool, comment ...*model.Comment) (*model.ArticleInfo) {
	return &model.ArticleInfo{
		ID: id, 
		Title: title,
		ImagePath: image,
		Nice: niceNum,
		Context: ctxt,
		UserStatus: &model.ArticleUserInfo{
			Nice: nice,
			List: list,
		},
		Comment: comment,
	}
}

func PackCooment(name, image, comment string) (*model.Comment) {
	return &model.Comment{
		Name: name,
		Image: image,
		Comment: comment,
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

func CastStringPointer(str string) *string {
	return &str
}
