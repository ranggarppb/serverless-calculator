package structs

import (
	"github.com/shopspring/decimal"
)

type CalculationInput struct {
	Input string `json:"input"`
}

type CalculationResult struct {
	Input  string `json:"input"`
	Result string `json:"result"`
}

type CalculationHistory struct {
	History []CalculationResult `json:"result"`
}

type CalculationWithOneInput struct {
	Input1    decimal.Decimal
	Operation string
}

type CalculationWithTwoInput struct {
	Input1    decimal.Decimal
	Input2    decimal.Decimal
	Operation string
}

type CalculationWithMultipleInput struct {
	Inputs []string
}
