package models

import "github.com/go-playground/validator/v10"

type TransactiontRequest struct {
	Reference string  `json:"reference" validate:"required"`
	Amount    float64 `json:"amount" validate:"required"`
}

func (t TransactiontRequest) Validate() error {
	v := validator.New()
	return v.Struct(t)
}

type BalanceResponse struct {
	Amount float64 `json:"balance"`
}
