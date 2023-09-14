package calculator

import (
	"context"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/shopspring/decimal"
)

type CalculatorInput struct {
	Input string `json:"input"`
}

type CalculatorResult struct {
	Input  string `json:"input"`
	Result string `json:"result"`
}

type CalculationHistory struct {
	History []CalculatorResult `json:"result"`
}

type CalculationWithOneInput struct {
	Input1    decimal.Decimal
	Operation string
}

type CalculationWithMultipleInput struct {
	Input1    decimal.Decimal
	Input2    decimal.Decimal
	Operation string
}

type ICalculatorService interface {
	GetCalculationHistory(context.Context) CalculationHistory
	Calculate(context.Context, string) (CalculatorResult, errors.WrappedError)
}
