package calculator

import (
	"strings"

	"github.com/ranggarppb/serverless-calculator/types"
	"github.com/ranggarppb/serverless-calculator/utils"
	"github.com/shopspring/decimal"
)

type calculatorService struct{}

func NewCalculatorService() *calculatorService {
	return &calculatorService{}
}

func (c *calculatorService) GetCalculationHistory() []string {
	result := []string{
		"Hello World!",
		"hello World!",
	}

	return result
}

func (c *calculatorService) Calculate(input string) (string, types.WrappedError) {

	parsedInput, err := c.parseInput(input)

	if err != nil {
		return "", err
	}

	switch parsedInput.Operation {
	case utils.ADDITION:
		return parsedInput.Input1.Add(parsedInput.Input2).String(), nil
	case utils.SUBSTRACTION:
		return parsedInput.Input1.Sub(parsedInput.Input2).String(), nil
	case utils.MULTIPLICATION:
		return parsedInput.Input1.Mul(parsedInput.Input2).String(), nil
	case utils.DIVISION:
		return parsedInput.Input1.Div(parsedInput.Input2).String(), nil
	default:
		return "", types.ErrInvalidOperation
	}
}

func (c *calculatorService) parseInput(input string) (types.CalculationWithTwoInput, types.WrappedError) {
	inputs := strings.Split(input, " ")

	if len(inputs) != 3 || !(utils.ContainString(utils.ALLOWED_OPERATIONS, inputs[1])) {
		return types.CalculationWithTwoInput{}, types.ErrInvalidOperation
	}

	num1, err := decimal.NewFromString(inputs[0])
	if err != nil {
		return types.CalculationWithTwoInput{}, types.ErrInvalidInputToBeOperated
	}
	num2, err := decimal.NewFromString(inputs[2])
	if err != nil {
		return types.CalculationWithTwoInput{}, types.ErrInvalidInputToBeOperated
	}

	return types.CalculationWithTwoInput{Input1: num1, Input2: num2, Operation: inputs[1]}, nil
}
