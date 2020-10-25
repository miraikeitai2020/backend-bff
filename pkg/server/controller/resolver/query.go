package resolver

import(
	"fmt"
	"time"
	"bytes"

	"context"
	"net/http"
	"io/ioutil"
	"github.com/miraikeitai2020/backend-bff/pkg/auth"
	"github.com/miraikeitai2020/backend-bff/pkg/utils"
	"github.com/miraikeitai2020/backend-bff/pkg/server/view"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

func (r *queryResolver) Signin(ctx context.Context) (*model.Token, error) {
	claims, errors := utils.ContextValueChecksum(ctx, "id", "pass")
	if len(errors) > 0 {
		return view.MakeSignResponse("", errors), nil
	}

	//TODO: Request Auth API
	id := claims["id"]
	
	token, err := auth.GenerateToken(id)
	if err != nil {
		errors = utils.MakeErrors(500, fmt.Sprintf("%v", err))
		return view.MakeSignResponse("", errors), nil
	}
	return view.MakeSignResponse(token, nil), nil
}

func (r *queryResolver) UserInfo(ctx context.Context) (*model.UserInfo, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeUserInfoResponse(nil, errors), nil
	}

	info := utils.PackUserInfo(
		"user name", 2020, 10, 25, 1,
		"euclid",
		"scp-jp",
		"如月工務店",
		"酩酊街",
	)
	return view.MakeUserInfoResponse(info, errors), nil
}

func (r *queryResolver) Like(ctx context.Context, articleid *string) (*model.Like, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeLikeResponse(false, errors), nil
	}

	return view.MakeLikeResponse(true, errors), nil
}

func (r *queryResolver) List(ctx context.Context) (*model.List, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeListResponse(nil, errors), nil
	}

	info := []*model.ArticleHeader {
		utils.PackArticleHeaderInfo(
			"SCP-1049-JP", "SCP-1049-JP", "http://scp-jp.wdfiles.com/local--files/scp-1049-jp/kakasi.jpg",
			"euclid",
			"scp-jp",
			"如月工務店",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1104-JP", "SCP-1104-JP", "http://scp-jp.wdfiles.com/local--files/scp-1104-jp/SCP-1104-JP.jpg",
			"safe",
			"scp-jp",
			"認識災害",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1955-JP", "SCP-1955-JP", "http://scp-jp.wdfiles.com/local--files/scp-1955-jp/Japanese-wolf.png",
			"euclid",
			"scp-jp",
			"犬",
			"酩酊街",
		),
	}
	return view.MakeListResponse(info, errors), nil
}

func (r *queryResolver) Genres(ctx context.Context) (*model.Genres, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeGenresResponse(errors), nil
	}

	return view.MakeGenresResponse(
		errors,
		"euclid",
		"scp-jp",
		"如月工務店",
		"酩酊街",
	), nil
}

func (r *queryResolver) Articles(ctx context.Context, genre string, year *int, month *int) (*model.Articles, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeArticlesResponse(nil, errors), nil
	}
	body, err := utils.MakeArticlesRequestJSON(genre, year, month)
	if err != nil {
		return view.MakeArticlesResponse(
			nil, utils.MakeErrors(
				500,
				fmt.Sprintf("%v", err),
			),
		), nil
	}
	request, err := http.NewRequest("GET", apiPath.Record + "/query/articles", bytes.NewBuffer(body))
	if err != nil {
		return view.MakeArticlesResponse(
			nil, utils.MakeErrors(
				500,
				fmt.Sprintf("%v", err),
			),
		), nil
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return view.MakeArticlesResponse(
			nil, utils.MakeErrors(
				500,
				fmt.Sprintf("%v", err),
			),
		), nil
	}
	body, err = ioutil.ReadAll(response.Body)
    if err != nil {
        return view.MakeArticlesResponse(
			nil, utils.MakeErrors(
				500,
				fmt.Sprintf("%v", err),
			),
		), nil
    }
	info := utils.MakeArticlesResponseStruct(body)
	return view.MakeArticlesResponse(info.ArticleList, errors), nil
}

func (r *queryResolver) ArticlesFromTag(ctx context.Context, tag string) (*model.Articles, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeArticlesResponse(nil, errors), nil
	}
	body, err := utils.MakeArticlesFromTagResponseJSON(tag)
	if err != nil {
		return view.MakeArticlesResponse(
			nil, utils.MakeErrors(500, fmt.Sprintf("%v", err))), nil
	}
	request, err := http.NewRequest("GET", apiPath.Record + "/query/tag/articles", bytes.NewBuffer(body))
	if err != nil {
		return view.MakeArticlesResponse(
			nil, utils.MakeErrors(500, fmt.Sprintf("%v", err))), nil
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return view.MakeArticlesResponse(
			nil, utils.MakeErrors(500, fmt.Sprintf("%v", err))), nil
	}
	body, err = ioutil.ReadAll(response.Body)
    if err != nil {
        return view.MakeArticlesResponse(nil, utils.MakeErrors(500, fmt.Sprintf("%v", err))), nil
    }
	info := utils.MakeArticlesResponseStruct(body)
	return view.MakeArticlesResponse(info.ArticleList, errors), nil
}

func (r *queryResolver) Article(ctx context.Context, articleid string) (*model.Article, error) {
	claims, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeArticleResponse(nil, errors), nil
	}
	m, err := auth.VerifyToken(claims["token"])
	if err != nil {
		return view.MakeArticleResponse(nil, utils.MakeErrors(500, fmt.Sprintf("%v", err))), nil
	}
	body, err := utils.MakeArticleRequestJSON(articleid)
	if err != nil {
		return view.MakeArticleResponse(nil, utils.MakeErrors(500, fmt.Sprintf("%v", err))), nil
	}
	request, err := http.NewRequest("GET", apiPath.Record + "/query/article", bytes.NewBuffer(body))
	if err != nil {
		return view.MakeArticleResponse(nil, utils.MakeErrors(500, fmt.Sprintf("%v", err))), nil
	}
	request.Header.Set("UserID", m["sub"].(string))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return view.MakeArticleResponse(nil, utils.MakeErrors(500,fmt.Sprintf("%v", err))), nil
	}
	body, err = ioutil.ReadAll(response.Body)
    if err != nil {
        return view.MakeArticleResponse(nil, utils.MakeErrors(500,fmt.Sprintf("%v", err))), nil
	}
	info := utils.MakeArticleResponseStruct(body)
	info.ID = articleid
	return view.MakeArticleResponse(&info, errors), nil
}

func (r *queryResolver) Log(ctx context.Context, logid string) (*model.Log, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeLogResponse(nil, errors), nil
	}

	info := utils.PackLogInfo(
		"1919810", "サービス残業", time.Now().String(), 810,
		[]float64{
			11.4,
			51.4,
			191.9,
			81.0,
		},
	)
	return view.MakeLogResponse(info, errors), nil
}

func (r *queryResolver) Logs(ctx context.Context) (*model.Logs, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeLogsResponse(nil, errors), nil
	}

	info := []*model.LogData{
		utils.PackLogData(
			"114514",
			"サービス残業",
		),
		utils.PackLogData(
			"1919",
			"サービス残業",
		),
		utils.PackLogData(
			"810",
			"サービス残業",
		),
	}
	return view.MakeLogsResponse(info, errors), nil
}

func (r *queryResolver) Spots(ctx context.Context, latitude float64, longitude float64, worktime int, emotion int) (*model.Spots, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeSpotsResponse(nil, nil, errors), nil
	}

	spot := utils.PackSpotInfo("9fdf8c2e", "五稜郭タワー", 41.7969245, 140.7545951)
	detour := []*model.Detour{
		utils.PackDetourInfo(
			"9fdf8c2e", "的場公園",
			"https://files.slack.com/files-pri/T013J36GEBZ-F01D9KA8SBC/matoba-park.jpg",
			"変わった滑り台や広い比較的広い砂場のような場所がある公園。",
			41.7795742, 140.7533145,
		),
	}
	return view.MakeSpotsResponse(spot, detour, errors), nil
}
