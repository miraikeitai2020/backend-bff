package view

import(
	errgen "errors"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/dto"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/service"
)

/*-------------------------------*/
/* Sign query response generator */
/*-------------------------------*/
func MakeSignResponse(token string, errors []*model.Errors) (*model.Token) {
	if len(errors) > 0 {
		return &model.Token {
			Value: "",
			Errors: errors,
		}
	}
	if token == "" {
		return &model.Token {
			Value: "",
			Errors: service.MakeErrors(500, errgen.New("Articles is empty.")),
		}
	}
	return &model.Token {
		Value: token,
		Errors: nil,
	}
}

/*-------------------------------*/
/*  Mutation response generator  */
/*-------------------------------*/
func MakeResultResponse(status bool, errors []*model.Errors) (*model.Result) {
	if len(errors) > 0 {
		return &model.Result {
			Status: false,
			Errors: errors,
		}
	}
	return &model.Result{
		Status: status,
		Errors: nil,
	}
}

/*-------------------------------*/
/*   Query response generator    */
/*-------------------------------*/
func MakeUserInfoResponse(info *model.User, errors []*model.Errors) (*model.UserInfo) {
	if len(errors) > 0 {
		return &model.UserInfo{
			Info: nil,
			Errors: errors,
		}
	}
	return &model.UserInfo{
		Info: info,
		Errors: nil,
	}
}

func MakeLikeResponse(status bool, errors []*model.Errors) (*model.Like) {
	if len(errors) > 0 {
		return &model.Like{
			Status: false,
			Errors: errors,
		}
	}
	return &model.Like{
		Status: status,
		Errors: nil,
	}
}

func MakeListResponse(list []*model.ArticleHeader, errors []*model.Errors) (*model.List) {
	if len(errors) > 0 {
		return &model.List{
			Articles: nil,
			Errors: errors,
		}
	}
	if len(list) == 0 {
		return &model.List{
			Articles: nil,
			Errors: service.MakeErrors(500, errgen.New("Articles is empty.")),
		}
	}
	return &model.List{
		Articles: list,
		Errors: errors,
	}
}

func MakeGenresResponse(errors []*model.Errors, genres ...string) (*model.Genres) {
	if len(errors) > 0 {
		return &model.Genres{
			Genres: nil,
			Errors: errors,
		}
	}
	if len(genres) == 0 {
		return &model.Genres{
			Genres: nil,
			Errors: service.MakeErrors(500, errgen.New("Articles is empty.")),
		}
	}
	return &model.Genres{
		Genres: genres,
		Errors: errors,
	}
}

func MakeArticlesResponse(body []byte, errors []*model.Errors) (*model.Articles) {
	info := service.ConvertResponseArticles(body)
	if len(errors) > 0 {
		return &model.Articles{
			Articles: nil,
			Errors: errors,
		}
	}
	if len(info.ArticleList) == 0 {
		return &model.Articles{
			Articles: nil,
			Errors: service.MakeErrors(500, errgen.New("Articles is empty.")),
		}
	}
	return &model.Articles{
		Articles: info.ArticleList,
		Errors: nil,
	}
}

func MakeArticleResponse(body []byte, errors []*model.Errors) (*model.Article) {
	info := service.ConvertResponseArticle(body)
	if len(errors) > 0 {
		return &model.Article{
			Info: nil,
			Errors: errors,
		}
	}
	return &model.Article{
		Info: &info.ArticleInfo,
		Errors: errors,
	}
}

func MakeLogResponse(info *model.LogInfo, errors []*model.Errors) (*model.Log) {
	if len(errors) > 0 {
		return &model.Log{
			Log: nil,
			Errors: errors,
		}
	}
	return &model.Log{
		Log: info,
		Errors: nil,
	}
}

func MakeLogsResponse(info []*model.LogData, errors []*model.Errors) (*model.Logs) {
	if len(errors) > 0 {
		return &model.Logs{
			Logs: nil,
			Errors: errors,
		}
	}
	if len(info) < 0 {
		return &model.Logs{
			Logs: nil,
			Errors: service.MakeErrors(500, errgen.New("Articles is empty.")),
		}
	}
	return &model.Logs{
		Logs: info,
		Errors: errors,
	}
}

func MakeSpotsResponse(spot *dto.SpotResponse, detour *dto.DetourResponse, errors []*model.Errors) (*model.Spots) {
	spots := service.ConvertSpotDtoToModel(*spot)
	detours := service.ConvertDetourDtoToModel(*detour)
	if len(errors) > 0 {
		return &model.Spots{
			Spot: nil,
			Detour: nil,
			Errors: errors,
		}
	}
	return &model.Spots{
		Spot: spots,
		Detour: detours,
		Errors: errors,
	}
}
