package types

import "github.com/shopspring/decimal"

type CalculatorInput struct {
	Input string `json:"input"`
}

type CalculatorResult struct {
	Input  string `json:"input"`
	Result string `json:"result"`
}

type CalculatorHistory struct {
	Result []string `json:"result"`
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
	GetCalculationHistory() []string
	Calculate(input string) (string, WrappedError)
}
