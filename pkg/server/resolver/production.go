package resolver

import (
	"time"
	"context"

	"github.com/miraikeitai2020/backend-bff/pkg/auth"
	"github.com/miraikeitai2020/backend-bff/pkg/utils"
	"github.com/miraikeitai2020/backend-bff/pkg/bff"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

func (r *mutationResolver) Signup(ctx context.Context) (*model.Token, error) {
	claims, errors := utils.ContextValueChecksum(ctx, "id", "pass")
	if len(errors) > 0 {
		return &model.Token{
			Value: "",
			Errors: errors,
		}, nil
	}

	//TODO: Request Auth API
	id := claims["id"]

	token, err := auth.GenerateToken(id)
	if err != nil {
	}
	return &model.Token{
		Value: token,
		Errors: []*model.Errors{},
	}, nil
}

func (r *mutationResolver) AddConstantUserInfo(ctx context.Context, gender int, year int, month int) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddName(ctx context.Context, name string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddGenre(ctx context.Context, genre []*string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddLike(ctx context.Context, articleid *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddList(ctx context.Context, articleid *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) DelList(ctx context.Context, articleid *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddRequest(ctx context.Context, genre *string, year *int, month *int, title *string, contents *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddComment(ctx context.Context, articleid *string, comment *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddNewLogData(ctx context.Context, date string, title string, worktime string, concentration string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddSubscription(ctx context.Context) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddEvaluation(ctx context.Context, spotid int, emotion int, value int) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *mutationResolver) AddSpot(ctx context.Context, name string, description string, image []int, latitude float64, longitude float64) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Result{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Result{
		Status: true,
		Errors: errors,
	}, nil
}

func (r *queryResolver) Signin(ctx context.Context) (*model.Token, error) {
	claims, errors := utils.ContextValueChecksum(ctx, "id", "pass")
	if len(errors) > 0 {
		return &model.Token{
			Value: "",
			Errors: errors,
		}, nil
	}

	//TODO: Request Auth API
	id := claims["id"]
	
	token, err := auth.GenerateToken(id)
	if err != nil {
	}
	return &model.Token{
		Value: token,
		Errors: []*model.Errors{},
	}, nil
}

func (r *queryResolver) UserInfo(ctx context.Context) (*model.UserInfo, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.UserInfo{
			Info: &model.User{},
			Errors: errors,
		}, nil
	}

	return &model.UserInfo{
		Info: &model.User{
			Name: "SaKu",
			Birthday: &model.Date{
				Year: 9,
				Month: 2,
			},
			Gender: 1,
			Genre: []*string {
				utils.CastStringPointer("アニメ・漫画"),
				utils.CastStringPointer("ほしまちすたじお"),
				utils.CastStringPointer("星街すいせい"),
				utils.CastStringPointer("vTuber"),
			},
		},
		Errors: nil,
	}, nil
}

func (r *queryResolver) Like(ctx context.Context, articleid *string) (*model.Like, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Like{
			Status: false,
			Errors: errors,
		}, nil
	}

	return &model.Like{
		Status: true,
		Errors: nil,
	}, nil
}

func (r *queryResolver) List(ctx context.Context, articleid *string) (*model.List, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.List{
			Articles: nil,
			Errors: errors,
		}, nil
	}

	return &model.List{
		Articles: []*model.ArticleHeader{
			&model.ArticleHeader{
				ID: "114514",
				Title: "50万人突破するまで歌い続ける",
				ImagePath: "https://pbs.twimg.com/media/Ei6rI-QVkAAGp5d?format=jpg&name=medium",
				Tags: []*string{
					utils.CastStringPointer("アニメ・漫画"),
					utils.CastStringPointer("ほしまちすたじお"),
					utils.CastStringPointer("星街すいせい"),
					utils.CastStringPointer("vTuber"),
				},
			},
		},
		Errors: nil,
	}, nil
}

func (r *queryResolver) Genres(ctx context.Context) (*model.Genres, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Genres{
			Genres: nil,
			Errors: errors,
		}, nil
	}

	return &model.Genres{
		Genres: []string{
			"アニメ・漫画",
			"映画",
			"声優",
			"海外",
		},
		Errors: nil,
	}, nil
}

func (r *queryResolver) Articles(ctx context.Context, genre string, year int, month int) (*model.Articles, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Articles{
			Articles: nil,
			Errors: errors,
		}, nil
	}

	return &model.Articles{
		Articles: []*model.ArticleHeader{
			&model.ArticleHeader{
				ID: "114514",
				Title: "50万人突破するまで歌い続ける",
				ImagePath: "https://pbs.twimg.com/media/Ei6rI-QVkAAGp5d?format=jpg&name=medium",
				Tags: []*string{
					utils.CastStringPointer("アニメ・漫画"),
					utils.CastStringPointer("ほしまちすたじお"),
					utils.CastStringPointer("星街すいせい"),
					utils.CastStringPointer("vTuber"),
				},
			},
		},
		Errors: nil,
	}, nil
}

func (r *queryResolver) Article(ctx context.Context, articleid string) (*model.Article, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Article{
			Info: nil,
			Errors: errors,
		}, nil
	}

	return &model.Article{
		Info: &model.ArticleInfo{
			Tile: "50万人突破するまで歌い続ける",
			ImagePath: "https://pbs.twimg.com/media/Ei6rI-QVkAAGp5d?format=jpg&name=medium",
			Nice: 1919810,
			Context: "星街すいせいとはvTuberである。2018年3月22日個人で活動するvTuberとしてデビュー、2019年5月19日にカバー株式会社の事務所であるホロライブプロダクションに所属し、現在は企業所属のvTuberとして活動している。2020/9/27の放送で彼女はチャンネル登録者が50万人に到達する瞬間を彼女の動画の視聴者と迎えようという趣旨の動画配信を行った。動画配信自体は歌配信メインで行われ、配信の途中で無事チャンネル登録者数が50万人に到達した。そして50万に到達した記念として50万人到達を祝う3Dモデルによるライブ配信をすることを発表した。ライブは2020年10月19日21:00から行われる予定だ。SNSでの盛り上がりは以下の盛り上がりを見せていた「もう50万達成してるやん、おめでとうございます ｱｱｧｧｧｧｱｱｱｧｧｧｧｱｱｱ!!!!」「Congratulations Suisei for the 500k subscribers !! I am super excited for your 3rd LIVE ! 50万おめでとう!!」「すいちゃん3Dライブやったぜぇぇぇ！前回すごく良かったから楽しみだ。すいちゃん登録50万おめでとう！」筆者としても彼女の活躍に大いに期待するところである。「今日もすいちゃんは可愛い！」",
			UserStatus: &model.ArticleUserInfo{
				Nice: true,
				List: true,
			},
			Comment: []*model.Comment{
				&model.Comment{
					Name: "SaKu",
					Image: "https://ca.slack-edge.com/T013J36GEBZ-U013S3FU8FQ-259987cfcb76-512",
					Comment: "今日もかわいいー！",
				},
				&model.Comment{
					Name: "takashi",
					Image: "https://ca.slack-edge.com/T013J36GEBZ-U013J3F33MM-6ce119c38fca-512",
					Comment: "今日もかわいいー！",
				},
			},
		},
		Errors: errors,
	}, nil
}

func (r *queryResolver) Log(ctx context.Context, logid int) (*model.Log, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Log{
			Log: nil,
			Errors: errors,
		}, nil
	}

	return &model.Log{
		Log: &model.LogInfo{
			ID: 1919810,
			Title: "サービス残業",
			Date: time.Now().String(),
			Worktime: 810,
			Concentration: []float64{
				11.4,
				51.4,
				191.9,
				81.0,
			},
		},
		Errors: nil,
	}, nil
}

func (r *queryResolver) Logs(ctx context.Context) (*model.Logs, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Logs{
			Logs: nil,
			Errors: errors,
		}, nil
	}

	return &model.Logs{
		Logs: []*model.LogData{
			&model.LogData{
				ID: 114514,
				Title: "サービス残業",
			},
			&model.LogData{
				ID: 1919,
				Title: "サービス残業",
			},
			&model.LogData{
				ID: 810,
				Title: "サービス残業",
			},
		},
		Errors: errors,
	}, nil
}

func (r *queryResolver) Spots(ctx context.Context, latitude float64, longitude float64, worktime int, emotion int) (*model.Spots, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return &model.Spots{
			Spots: nil,
			Errors: errors,
		}, nil
	}

	return &model.Spots{
		Spots: []*model.Spot{
			&model.Spot{
				ID: 114,
				Name: "与那国島",
				Image: "https://upload.wikimedia.org/wikipedia/commons/9/99/Yonaguni_Island_ISS039.jpg",
				Description: "南西諸島八重山列島の島。島としては日本の最西端に位置し、国境の島と呼ばれる。",
				Locate: &model.Locate{
					Latitude: 24.2705,
    				Longitude: 122.5557,
				},
			},
			&model.Spot{
				ID: 514,
				Name: "南鳥島",
				Image: "https://upload.wikimedia.org/wikipedia/commons/d/d2/Aerial-View-Minamitori-Island-1987.jpg",
				Description: "小笠原諸島の島。本州から1,800 km離れた日本の最東端としても知られている。",
				Locate: &model.Locate{
					Latitude: 24.1659,
    				Longitude: 153.5912,
				},
			},
			&model.Spot{
				ID: 1919,
				Name: "沖ノ鳥島",
				Image: "https://upload.wikimedia.org/wikipedia/commons/f/f2/Okinotorishima20070602.jpg",
				Description: "太平洋（フィリピン海）上に位置する小笠原諸島に属する孤立島",
				Locate: &model.Locate{
					Latitude: 20.2531,
    				Longitude: 136.0411,
				},
			},
			&model.Spot{
				ID: 810,
				Name: "択捉島",
				Image: "https://upload.wikimedia.org/wikipedia/commons/4/43/Iturup_Terra_cut.jpg",
				Description: "千島列島南部に位置する同列島内で面積が最大の島。複数の活火山が存在する火山島である。",
				Locate: &model.Locate{
					Latitude: 45.3326,
    				Longitude: 148.4508,
				},
			},
		},
		Errors: errors,
	}, nil
}

// Mutation returns bff.MutationResolver implementation.
func (r *Resolver) Mutation() bff.MutationResolver { return &mutationResolver{r} }

// Query returns bff.QueryResolver implementation.
func (r *Resolver) Query() bff.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
