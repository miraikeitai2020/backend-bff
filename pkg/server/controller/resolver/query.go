package resolver

import (
	"context"

	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/dao"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/service"
	"github.com/miraikeitai2020/backend-bff/pkg/server/view"
	"github.com/miraikeitai2020/backend-bff/pkg/utils"
)

func (r *queryResolver) Signin(ctx context.Context) (*model.Token, error) {
	claims, errors := service.CreateHeaderLoader(ctx, "id", "pass")
	if len(errors) > 0 {
		return view.MakeSignResponse("", errors), nil
	}

	//TODO: Request Auth API
	id := claims["id"]

	token, err := service.GenerateToken(id)
	if err != nil {
		return view.MakeSignResponse("", service.MakeErrors(500, err)), nil
	}
	return view.MakeSignResponse(token, nil), nil
}

func (r *queryResolver) UserInfo(ctx context.Context) (*model.UserInfo, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeUserInfoResponse(nil, errors), nil
	}

	info := utils.PackUserInfo(
		"user name", 2020, 10, 25, 1,
		"ドラマ",
		"音楽",
		"その他",
		"政治",
	)
	return view.MakeUserInfoResponse(info, errors), nil
}

func (r *queryResolver) Like(ctx context.Context, articleid *string) (*model.Like, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeLikeResponse(false, errors), nil
	}

	return view.MakeLikeResponse(true, errors), nil
}

func (r *queryResolver) List(ctx context.Context) (*model.List, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeListResponse(nil, errors), nil
	}

	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeListResponse(nil, service.MakeErrors(500, err)), nil
	}

	client := dao.MakeListClient()
	body, err := client.Request(claims.Load("sub"))
	if err != nil {
		return view.MakeListResponse(nil, service.MakeErrors(500, err)), nil
	}
	info := service.ConvertResponseList(body)

	return view.MakeListResponse(info.Articles, errors), nil
}

func (r *queryResolver) Genres(ctx context.Context) (*model.Genres, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeGenresResponse(errors), nil
	}

	return view.MakeGenresResponse(
		errors,
		"ドラマ",
		"音楽",
		"その他",
		"政治",
	), nil
}

func (r *queryResolver) Articles(ctx context.Context, genre string, year *int, month *int) (*model.Articles, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeArticlesResponse(nil, errors), nil
	}

	// Request Record API //
	client := dao.MakeArticlesClient(genre, year, month)
	body, err := client.Request()
	if err != nil {
		return view.MakeArticlesResponse(nil, service.MakeErrors(500, err)), nil
	}

	// Retrun Response //
	return view.MakeArticlesResponse(body, errors), nil
}

func (r *queryResolver) ArticlesFromTag(ctx context.Context, tag string) (*model.Articles, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeArticlesResponse(nil, errors), nil
	}

	// Request Record API //
	client := dao.MakeArticlesFromTagClient(tag)
	body, err := client.Request()
	if err != nil {
		return view.MakeArticlesResponse(nil, service.MakeErrors(500, err)), nil
	}

	// Retrun Response //
	return view.MakeArticlesResponse(body, errors), nil
}

func (r *queryResolver) Article(ctx context.Context, articleid string) (*model.Article, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeArticleResponse(nil, errors), nil
	}
	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeArticleResponse(nil, service.MakeErrors(500, err)), nil
	}

	client := dao.MakeArticleClient(articleid)
	body, err := client.Request(claims.Load("sub"))
	if err != nil {
		return view.MakeArticleResponse(nil, service.MakeErrors(500, err)), nil
	}

	return view.MakeArticleResponse(body, errors), nil
}

func (r *queryResolver) Log(ctx context.Context, logid string) (*model.Log, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeLogResponse(nil, errors), nil
	}

	client := dao.NewLogClient()
	body, err := client.Request(logid)
	if err != nil {
		return view.MakeLogResponse(nil, service.MakeErrors(500, err)), nil
	}

	log := service.ConvertResponseLog(body)

	return view.MakeLogResponse(&log.Info, errors), nil
}

func (r *queryResolver) Logs(ctx context.Context) (*model.Logs, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeLogsResponse(nil, errors), nil
	}

	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeLogsResponse(nil, service.MakeErrors(500, err)), nil
	}

	client := dao.NewLogsClient()
	body, err := client.Request(claims.Load("sub"))
	if err != nil {
		return view.MakeLogsResponse(nil, service.MakeErrors(500, err)), nil
	}

	info := service.ConvertResponseLogs(body)

	return view.MakeLogsResponse(info.Data, errors), nil
}

func (r *queryResolver) Spots(ctx context.Context, latitude float64, longitude float64, worktime int, emotion int) (*model.Spots, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeSpotsResponse(nil, nil, errors), nil
	}
	client := dao.MakeSpotClient(latitude, longitude, worktime, emotion)
	body, err := client.Request()
	if err != nil {
		return view.MakeSpotsResponse(nil, nil, service.MakeErrors(500, err)), nil
	}
	spot := service.ConvertResponseSpot(body)

	//
	client = dao.MakeDetourClient(spot.Spot.Latitude, spot.Spot.Longitude, latitude, longitude, worktime, emotion)
	body, err = client.Request()
	if err != nil {
		return view.MakeSpotsResponse(nil, nil, service.MakeErrors(500, err)), nil
	}

	detour := service.ConvertResponseDetour(body)

	return view.MakeSpotsResponse(&spot, &detour, errors), nil
}
