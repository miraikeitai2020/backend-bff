package service

import(
	"fmt"
	"context"
	"github.com/miraikeitai2020/backend-bff/pkg/server/model"
)

func CreateHeaderLoader(ctx context.Context, keys ...string) (claims map[string]string, errors []*model.Errors) {
	if len(keys) < 1 {
		return
	}

	var value string
	var ok bool
	claims = make(map[string]string)

	for _, key := range keys {
		if value, ok = ctx.Value(key).(string); !ok {
			errors = append(errors, MakeError(500, fmt.Sprintf("Faild get HTTP header value: %s", key)))
		}
		if value == "" {
			errors = append(errors, MakeError(400, fmt.Sprintf("HTTP header value is empty: %s", key)))
		}
		claims[key] = value
	}
	return
}
