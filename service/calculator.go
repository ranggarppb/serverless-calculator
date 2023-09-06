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

	switch input := parsedInput.(type) {
	case types.CalculationWithOneInput:
		return c.doCalculationWithOneInput(input)
	case types.CalculationWithMultipleInput:
		return c.doCalculationWithMultipleInput(input)
	default:
		return "", types.ErrInvalidOperation
	}
}

func (c *calculatorService) parseInput(input string) (interface{}, types.WrappedError) {
	inputs := strings.Split(input, " ")

	calculationType, err := c.getCalculationType(inputs)

	if err != nil {
		return nil, err
	}

	switch calculationType.(type) {
	case types.CalculationWithOneInput:
		res, err := c.validateAndConstructCalculationOneInput(inputs)

		if err != nil {
			return types.CalculationWithOneInput{}, err
		}

		return res, nil
	case types.CalculationWithMultipleInput:
		res, err := c.validateAndConstructCalculationMultipleInput(inputs)

		if err != nil {
			return types.CalculationWithMultipleInput{}, err
		}

		return res, nil
	default:
		return nil, types.ErrInvalidOperation
	}

}

func (c *calculatorService) getCalculationType(inputs []string) (interface{}, types.WrappedError) {
	if len(inputs) == 0 {
		return nil, types.ErrInvalidInput
	}

	_, err := decimal.NewFromString(inputs[0])

	if err != nil {
		return types.CalculationWithOneInput{}, nil
	} else {
		return types.CalculationWithMultipleInput{}, nil
	}
}

func (c *calculatorService) validateAndConstructCalculationOneInput(inputs []string) (types.CalculationWithOneInput, types.WrappedError) {
	if len(inputs) != 2 || !(utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, inputs[0])) {
		return types.CalculationWithOneInput{}, types.ErrInvalidOperation
	}

	num, err := decimal.NewFromString(inputs[1])

	if err != nil {
		return types.CalculationWithOneInput{}, types.ErrInvalidInputToBeOperated
	}
	return types.CalculationWithOneInput{Input1: num, Operation: inputs[0]}, nil
}

func (c *calculatorService) doCalculationWithOneInput(input types.CalculationWithOneInput) (string, types.WrappedError) {
	switch input.Operation {
	case utils.NEGATION:
		return input.Input1.Neg().String(), nil
	case utils.SQUARE:
		power := decimal.NewFromInt(2)
		return input.Input1.Pow(power).String(), nil
	case utils.SQUAREROOT:
		power := decimal.NewFromFloat(0.5)
		return input.Input1.Pow(power).String(), nil
	case utils.ABSOLUTE:
		return input.Input1.Abs().String(), nil
	default:
		return "", types.ErrInvalidOperation
	}
}

func (c *calculatorService) validateAndConstructCalculationMultipleInput(inputs []string) (types.CalculationWithMultipleInput, types.WrappedError) {
	if len(inputs) != 3 || !(utils.ContainString(utils.OPERATIONS_WITH_TWO_INPUTS, inputs[1])) {
		return types.CalculationWithMultipleInput{}, types.ErrInvalidOperation
	}

	num1, err1 := decimal.NewFromString(inputs[0])
	num2, err2 := decimal.NewFromString(inputs[2])

	if !(err1 == nil && err2 == nil) {
		return types.CalculationWithMultipleInput{}, types.ErrInvalidInputToBeOperated
	}
	return types.CalculationWithMultipleInput{Input1: num1, Input2: num2, Operation: inputs[1]}, nil
}

func (c *calculatorService) doCalculationWithMultipleInput(input types.CalculationWithMultipleInput) (string, types.WrappedError) {
	switch input.Operation {
	case utils.ADDITION:
		return input.Input1.Add(input.Input2).String(), nil
	case utils.SUBSTRACTION:
		return input.Input1.Sub(input.Input2).String(), nil
	case utils.MULTIPLICATION:
		return input.Input1.Mul(input.Input2).String(), nil
	case utils.DIVISION:
		return input.Input1.Div(input.Input2).String(), nil
	default:
		return "", types.ErrInvalidOperation
	}
}
