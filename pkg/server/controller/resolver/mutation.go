package resolver

import(
	"fmt"
	"context"
	"github.com/miraikeitai2020/backend-bff/pkg/auth"
	"github.com/miraikeitai2020/backend-bff/pkg/utils"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
	"github.com/miraikeitai2020/backend-bff/pkg/server/view"
)

func (r *mutationResolver) Signup(ctx context.Context) (*model.Token, error) {
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

func (r *mutationResolver) AddConstantUserInfo(ctx context.Context, gender int, year int, month int, date int) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddName(ctx context.Context, name string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddGenre(ctx context.Context, genre []*string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddLike(ctx context.Context, articleid *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddList(ctx context.Context, articleid *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) DelList(ctx context.Context, articleid *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddRequest(ctx context.Context, genre *string, year *int, month *int, title *string, contents *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddComment(ctx context.Context, articleid *string, comment *string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddNewLogData(ctx context.Context, id string, date string, title string, worktime string, concentration string) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddSubscription(ctx context.Context) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddEvaluation(ctx context.Context, spotid string, emotion int, status bool) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddSpot(ctx context.Context, name string, description string, image string, latitude float64, longitude float64) (*model.Result, error) {
	_, errors := utils.ContextValueChecksum(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}
