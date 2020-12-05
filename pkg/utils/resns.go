package utils

import "github.com/miraikeitai2020/backend-bff/pkg/server/model"

func PackNiceData(nice *int) (*model.Nice) {
	return &model.Nice{
		Nice: *nice,
	}
}
