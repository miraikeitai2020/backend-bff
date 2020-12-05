package resolver

import (
	"context"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/dao"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model/service"
	"github.com/miraikeitai2020/backend-bff/pkg/server/view"
)

func (r *mutationResolver) Signup(ctx context.Context) (*model.Token, error) {
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

func (r *mutationResolver) AddConstantUserInfo(ctx context.Context, gender int, year int, month int, date int) (*model.Result, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddName(ctx context.Context, name string) (*model.Result, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddGenre(ctx context.Context, genre []*string) (*model.Result, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddNice(ctx context.Context, articleid *string) (*model.Nice, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeNiceResponse(nil, errors), nil
	}
	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeNiceResponse(nil, service.MakeErrors(500, err)), nil
	}
	client := dao.MakeAddNiceClient(*articleid)
	body, err := client.Request(claims.Load("sub"))
	if err != nil {
		return view.MakeNiceResponse(nil, service.MakeErrors(500, err)), nil
	}
	info := service.ConvertResponseNice(body)
	return view.MakeNiceResponse(&info.NiceInfo, errors), nil
}

func (r *mutationResolver) AddList(ctx context.Context, articleid *string) (*model.Result, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}
	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}
	client := dao.MakeAddListClient(*articleid)
	if _, err = client.Request(claims.Load("sub")); err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) DelList(ctx context.Context, articleid *string) (*model.Result, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}
	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}
	client := dao.MakeDelListClient(*articleid)
	if _, err = client.Request(claims.Load("sub")); err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddRequest(ctx context.Context, genre *string, year *int, month *int, title *string, contents *string) (*model.Result, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	client := dao.MakeAddRequestClient(*genre, *year, *month, *title, *contents)
	if _, err = client.Request(claims.Load("sub")); err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddComment(ctx context.Context, articleid *string, comment *string) (*model.Result, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	client := dao.MakeAddCommentClient(*articleid, *comment)
	if _, err = client.Request(claims.Load("sub")); err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddNewLogData(ctx context.Context, id string, date string, title string, worktime int, concentration []float64) (*model.Result, error) {
	header, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	claims, err := service.VerifyToken(header["token"])
	if err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	client := dao.NewAddLogClient(title, date, worktime, concentration)
	if _, err = client.Request(claims.Load("sub")); err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddSubscription(ctx context.Context) (*model.Result, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}

	return view.MakeResultResponse(true, errors), nil
}

func (r *mutationResolver) AddEvaluation(ctx context.Context, spotid string, emotion int, status bool) (*model.Result, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}
	client := dao.MakeEvaluationClient(spotid, emotion, status)
	body, err := client.Request()
	if err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}
	info := service.ConvertResponseMutation(body)

	return view.MakeResultResponse(info.Status, errors), nil
}

func (r *mutationResolver) AddSpot(ctx context.Context, name string, description string, image string, latitude float64, longitude float64) (*model.Result, error) {
	_, errors := service.CreateHeaderLoader(ctx, "token")
	if len(errors) > 0 {
		return view.MakeResultResponse(false, errors), nil
	}
	client := dao.MakeAddSpotClient(name, image, description, latitude, longitude)
	body, err := client.Request()
	if err != nil {
		return view.MakeResultResponse(false, service.MakeErrors(500, err)), nil
	}
	info := service.ConvertResponseMutation(body)
	return view.MakeResultResponse(info.Status, errors), nil
}
