package resolver

import (
	"context"

	"github.com/miraikeitai2020/backend-bff-API/pkg/bff"
	"github.com/miraikeitai2020/backend-bff-API/pkg/server/model"
)

func (r *mutationResolver) AddArticle(ctx context.Context, index *int, article *string) (*model.Article, error) {
	return &model.Article{
		Index: *index,
		Article: *article,
	}, nil
}

func (r *mutationResolver) PostUserData(ctx context.Context, emotion float64, city string) (*model.CityData, error) {
	if emotion >= 0 && emotion < 0.3 {
		return &model.CityData{
			Name: "沖縄県　与那国島",
			Latitude: 24.2705,
			Longitude: 122.5557,
		}, nil
	}
	if emotion >= 0.3 && emotion < 0.6 {
		return &model.CityData{
			Name: "東京都　南鳥島",
			Latitude: 24.1659,
			Longitude: 153.5912,
		}, nil
	}
	if emotion >= 0.6 && emotion < 0.9 {
		return &model.CityData{
			Name: "東京都　沖ノ鳥島",
			Latitude: 20.2531,
			Longitude: 136.0411,
		}, nil
	}
	return &model.CityData{
		Name: "北海道　択捉島",
		Latitude: 45.3326,
		Longitude: 148.4508,
	}, nil
}

func (r *queryResolver) Odd(ctx context.Context, number int) (*model.IsOdd, error) {
	if number%2 != 1 {
		return &model.IsOdd {
			Number: number,
			Judge: false,
		}, nil
	}
	return &model.IsOdd {
		Number: number,
		Judge: true,
	}, nil
}

func (r *queryResolver) Even(ctx context.Context, number int) (*model.IsEven, error) {
	if number%2 != 0 {
		return &model.IsEven {
			Number: number,
			Judge: false,
		}, nil
	}
	return &model.IsEven {
		Number: number,
		Judge: true,
	}, nil
}

func (r *queryResolver) City(ctx context.Context, id *int) (*model.CityData, error) {
	if *id == 1 {
		return &model.CityData{
			Name: "沖縄県　与那国島",
			Latitude: 24.2705,
			Longitude: 122.5557,
		}, nil
	}
	if *id == 2 {
		return &model.CityData{
			Name: "東京都　南鳥島",
			Latitude: 24.1659,
			Longitude: 153.5912,
		}, nil
	}
	if *id == 3 {
		return &model.CityData{
			Name: "東京都　沖ノ鳥島",
			Latitude: 20.2531,
			Longitude: 136.0411,
		}, nil
	}
	return &model.CityData{
		Name: "北海道　択捉島",
		Latitude: 45.3326,
		Longitude: 148.4508,
	}, nil
}

func (r *queryResolver) AllCity(ctx context.Context) ([]*model.CityData, error) {
	return []*model.CityData{
		&model.CityData{
			Name: "沖縄県　与那国島",
			Latitude: 24.2705,
			Longitude: 122.5557,
		},
		&model.CityData{
			Name: "東京都　沖ノ鳥島",
			Latitude: 20.2531,
			Longitude: 136.0411,
		},
		&model.CityData{
			Name: "東京都　沖ノ鳥島",
			Latitude: 20.2531,
			Longitude: 136.0411,
		},
		&model.CityData{
			Name: "東京都　南鳥島",
			Latitude: 24.1659,
			Longitude: 153.5912,
		},
	}, nil
}

// Mutation returns bff.MutationResolver implementation.
func (r *Resolver) Mutation() bff.MutationResolver { return &mutationResolver{r} }

// Query returns bff.QueryResolver implementation.
func (r *Resolver) Query() bff.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
