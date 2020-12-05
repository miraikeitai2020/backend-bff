package dto

import "github.com/miraikeitai2020/backend-bff/pkg/server/model"

type NiceResponse struct {
	NiceInfo model.Nice `json:"nice"`
}

type MutationResponse struct {
	Status	bool	`json:"status"`
}

