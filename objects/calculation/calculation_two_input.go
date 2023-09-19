package calculation

import (
	"fmt"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/utils"
	"github.com/shopspring/decimal"
)

type calculationWithTwoInput struct {
	Input1    string
	Input2    string
	Operation string
}

func NewCalculationWithTwoInput(input1 string, input2 string, operation string) *calculationWithTwoInput {
	return &calculationWithTwoInput{Input1: input1, Input2: input2, Operation: operation}
}

func (i calculationWithTwoInput) GetInput() string {
	return fmt.Sprintf("%s %s %s", i.Input1, i.Operation, i.Input2)
}

func (i calculationWithTwoInput) Validate() errors.WrappedError {
	if !utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, i.Operation) {
		return errors.ErrInvalidOperation
	}

	_, err1 := decimal.NewFromString(i.Input1)
	_, err2 := decimal.NewFromString(i.Input2)

	if err1 != nil || err2 != nil {
		return errors.ErrInvalidInputToBeOperated
	}

	return nil
}

func (i calculationWithTwoInput) Calculate() string {
	input1Dec, _ := decimal.NewFromString(i.Input1)
	input2Dec, _ := decimal.NewFromString(i.Input2)
	switch i.Operation {
	case utils.ADDITION:
		return input1Dec.Add(input2Dec).String()
	case utils.SUBSTRACTION:
		return input1Dec.Sub(input2Dec).String()
	case utils.MULTIPLICATION:
		return input1Dec.Mul(input2Dec).String()
	case utils.DIVISION:
		return input1Dec.Div(input2Dec).String()
	default:
		return ""
	}
}
