package utils

import(
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

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

func CastStringPointer(str string) *string {
	return &str
}
