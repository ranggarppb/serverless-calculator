package calculator

import (
	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/shopspring/decimal"
)

type CalculationInput struct {
	Input string `json:"input"`
}

func (i *CalculationInput) GetCalculationType() interface{} {
	return nil
}

type CalculationWithOneInput struct {
	Input1    decimal.Decimal
	Operation string
}

func (i *CalculationWithOneInput) Validate() errors.WrappedError {
	return nil
}

type CalculationWithMultipleInput struct {
	Inputs []string
}

func (i *CalculationWithMultipleInput) Validate() errors.WrappedError {
	return nil
}

type CalculationWithTwoInput struct {
	Input1    decimal.Decimal
	Input2    decimal.Decimal
	Operation string
}
