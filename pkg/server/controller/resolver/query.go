package resolver

import(
	"fmt"
	"time"

	"context"
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
			"SCP-1049-JP", "SCP-1049-JP", "image URL",
			"euclid",
			"scp-jp",
			"如月工務店",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1104-JP", "SCP-1104-JP", "image URL",
			"safe",
			"scp-jp",
			"認識災害",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1955-JP", "SCP-1955-JP", "image URL",
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
		return &model.Articles{
			Articles: nil,
			Errors: errors,
		}, nil
	}
	if year == nil && month != nil {
		return view.MakeArticlesResponse(
			nil,
			utils.MakeErrors(400, "Argment `year` is emptly"),
		), nil
	}
	if year != nil && month == nil {
		return view.MakeArticlesResponse(
			nil,
			utils.MakeErrors(400, "Argment `year` is emptly"),
		), nil
	}

	info := []*model.ArticleHeader {
		utils.PackArticleHeaderInfo(
			"SCP-1049-JP", "SCP-1049-JP", "image URL",
			"euclid",
			"scp-jp",
			"如月工務店",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1104-JP", "SCP-1104-JP", "image URL",
			"safe",
			"scp-jp",
			"認識災害",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1955-JP", "SCP-1955-JP", "image URL",
			"euclid",
			"scp-jp",
			"犬",
			"酩酊街",
		),
	}
	return view.MakeArticlesResponse(info, errors), nil
}

func (r *queryResolver) ArticlesFromTag(ctx context.Context, tag string) (*model.Articles, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Articles{
			Articles: nil,
			Errors: errors,
		}, nil
	}

	info := []*model.ArticleHeader {
		utils.PackArticleHeaderInfo(
			"SCP-1049-JP", "SCP-1049-JP", "image URL",
			"euclid",
			"scp-jp",
			"如月工務店",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1104-JP", "SCP-1104-JP", "image URL",
			"safe",
			"scp-jp",
			"認識災害",
			"酩酊街",
		),
		utils.PackArticleHeaderInfo(
			"SCP-1955-JP", "SCP-1955-JP", "image URL",
			"euclid",
			"scp-jp",
			"犬",
			"酩酊街",
		),
	}
	return view.MakeArticlesResponse(info, errors), nil
}

func (r *queryResolver) Article(ctx context.Context, articleid string) (*model.Article, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeArticleResponse(nil, errors), nil
	}

	info := utils.PackArticleInfo(
		"50万人突破するまで歌い続ける", "image URL", 1919810,
		"星街すいせいとはvTuberである。2018年3月22日個人で活動するvTuberとしてデビュー、2019年5月19日にカバー株式会社の事務所であるホロライブプロダクションに所属し、現在は企業所属のvTuberとして活動している。2020/9/27の放送で彼女はチャンネル登録者が50万人に到達する瞬間を彼女の動画の視聴者と迎えようという趣旨の動画配信を行った。動画配信自体は歌配信メインで行われ、配信の途中で無事チャンネル登録者数が50万人に到達した。そして50万に到達した記念として50万人到達を祝う3Dモデルによるライブ配信をすることを発表した。ライブは2020年10月19日21:00から行われる予定だ。SNSでの盛り上がりは以下の盛り上がりを見せていた「もう50万達成してるやん、おめでとうございます ｱｱｧｧｧｧｱｱｱｧｧｧｧｱｱｱ!!!!」「Congratulations Suisei for the 500k subscribers !! I am super excited for your 3rd LIVE ! 50万おめでとう!!」「すいちゃん3Dライブやったぜぇぇぇ！前回すごく良かったから楽しみだ。すいちゃん登録50万おめでとう！」筆者としても彼女の活躍に大いに期待するところである。「今日もすいちゃんは可愛い！」",
		true, true,
		utils.PackCooment(
			"User Name",
			"https://ca.slack-edge.com/T013J36GEBZ-U013S3FU8FQ-259987cfcb76-512",
			"今日もかわいいー！",
		),
		utils.PackCooment(
			"User Name",
			"https://ca.slack-edge.com/T013J36GEBZ-U013J3F33MM-6ce119c38fca-512",
			"今日もかわいいー！",
		),
	)
	return view.MakeArticleResponse(info, errors), nil
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
