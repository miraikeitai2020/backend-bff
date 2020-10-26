package service

import(
	"fmt"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

const(
	STATUS_CODE_400 = "Bad Request"
	STATUS_CODE_403 = "Forbidden"
	STATUS_CODE_500 = "Internal Server Error"
)

func MakeError(code int, msg string) (errors *model.Errors) {
	switch code {
	case 400:
		errors = &model.Errors{
			Code: 400,
			Message: STATUS_CODE_400,
			Description: msg,
		}
	case 403:
		errors = &model.Errors{
			Code: 403,
			Message: STATUS_CODE_403,
			Description: msg,
		}
	case 500:
		errors = &model.Errors{
			Code: 500,
			Message: STATUS_CODE_500,
			Description: msg,
		}
	}
	return
}

func MakeErrors(code int, err error) ([]*model.Errors) {
	return []*model.Errors{
		MakeError(code, fmt.Sprintf("%v", err)),
	}
}
