package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"github.com/miraikeitai2020/backend-bff/pkg/bff"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

type Resolver struct{}

func (r *mutationResolver) Signup(ctx context.Context, id string, password string) (*model.Token, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddConstantUserInfo(ctx context.Context, token string, gender int, year int, month int) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddName(ctx context.Context, token string, name string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) Addgenre(ctx context.Context, token string, genre []*string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddLike(ctx context.Context, userid *string, articleid *string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddList(ctx context.Context, userid *string, articleid *string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) DelList(ctx context.Context, userid *string, articleid *string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddComment(ctx context.Context, token *string, id *string, comment *string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddNewLogData(ctx context.Context, id string, date string, title string, worktime string, concentration string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddSubscription(ctx context.Context, id string) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddEvaluation(ctx context.Context, id int, emotion int, value int) (*model.Result, error) {
	panic("not implemented")
}

func (r *mutationResolver) AddSpot(ctx context.Context, name string, description string, image []int, latitude float64, longitude float64) (*model.Result, error) {
	panic("not implemented")
}

func (r *queryResolver) Signin(ctx context.Context, id string, password string) (*model.Token, error) {
	panic("not implemented")
}

func (r *queryResolver) UserInfo(ctx context.Context, token string) (*model.UserInfo, error) {
	panic("not implemented")
}

func (r *queryResolver) Like(ctx context.Context, userid *string, articleid *string) (*model.Like, error) {
	panic("not implemented")
}

func (r *queryResolver) List(ctx context.Context, userid *string, articleid *string) (*model.List, error) {
	panic("not implemented")
}

func (r *queryResolver) Genres(ctx context.Context) (*model.Genres, error) {
	panic("not implemented")
}

func (r *queryResolver) Articles(ctx context.Context, genre string, year int, month int) (*model.Articles, error) {
	panic("not implemented")
}

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	panic("not implemented")
}

func (r *queryResolver) Log(ctx context.Context, id string, logID string) (*model.Log, error) {
	panic("not implemented")
}

func (r *queryResolver) Logs(ctx context.Context, id *string) (*model.Logs, error) {
	panic("not implemented")
}

func (r *queryResolver) Spots(ctx context.Context, latitude float64, longitude float64, worktime int, emotion int) (*model.Spots, error) {
	panic("not implemented")
}

// Mutation returns bff.MutationResolver implementation.
func (r *Resolver) Mutation() bff.MutationResolver { return &mutationResolver{r} }

// Query returns bff.QueryResolver implementation.
func (r *Resolver) Query() bff.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
