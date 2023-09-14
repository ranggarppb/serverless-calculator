package calculator

import (
	"context"
	"math"
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/utils"
	"github.com/shopspring/decimal"
)

type calculatorService struct{}

func NewCalculatorService() *calculatorService {
	return &calculatorService{}
}

func (c *calculatorService) GetCalculationHistory(ctx context.Context) CalculationHistory {

	return CalculationHistory{}
}

func (c *calculatorService) Calculate(ctx context.Context, input string) (CalculatorResult, errors.WrappedError) {

	parsedInput, err := c.parseInput(input)

	if err != nil {
		return CalculatorResult{}, err
	}

	switch inputType := parsedInput.(type) {
	case CalculationWithOneInput:
		result, err := c.doCalculationWithOneInput(inputType)
		if err != nil {
			return CalculatorResult{}, err
		}
		return CalculatorResult{Input: input, Result: result}, nil
	case CalculationWithMultipleInput:
		result, err := c.doCalculationWithMultipleInput(inputType)
		if err != nil {
			return CalculatorResult{}, err
		}
		return CalculatorResult{Input: input, Result: result}, nil
	default:
		return CalculatorResult{}, errors.ErrInvalidOperation
	}
}

func (c *calculatorService) parseInput(input string) (interface{}, errors.WrappedError) {
	inputs := strings.Split(input, " ")

	calculationType, err := c.getCalculationType(inputs)

	if err != nil {
		return nil, err
	}

	switch calculationType.(type) {
	case CalculationWithOneInput:
		res, err := c.validateAndConstructCalculationOneInput(inputs)

		if err != nil {
			return CalculationWithOneInput{}, err
		}

		return res, nil
	case CalculationWithMultipleInput:
		res, err := c.validateAndConstructCalculationMultipleInput(inputs)

		if err != nil {
			return CalculationWithMultipleInput{}, err
		}

		return res, nil
	default:
		return nil, errors.ErrInvalidOperation
	}

}

func (c *calculatorService) getCalculationType(inputs []string) (interface{}, errors.WrappedError) {
	if len(inputs) == 0 {
		return nil, errors.ErrInvalidInput
	}

	_, err := decimal.NewFromString(inputs[0])

	if err != nil {
		return CalculationWithOneInput{}, nil
	} else {
		return CalculationWithMultipleInput{}, nil
	}
}

func (c *calculatorService) validateAndConstructCalculationOneInput(inputs []string) (CalculationWithOneInput, errors.WrappedError) {
	if len(inputs) != 2 || !(utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, inputs[0])) {
		return CalculationWithOneInput{}, errors.ErrInvalidOperation
	}

	num, err := decimal.NewFromString(inputs[1])

	if err != nil {
		return CalculationWithOneInput{}, errors.ErrInvalidInputToBeOperated
	}
	if inputs[0] == utils.SQUAREROOT && num.LessThan(decimal.Zero) {
		return CalculationWithOneInput{}, errors.ErrInvalidInputToBeOperated
	}
	return CalculationWithOneInput{Input1: num, Operation: inputs[0]}, nil
}

func (c *calculatorService) doCalculationWithOneInput(input CalculationWithOneInput) (string, errors.WrappedError) {
	switch input.Operation {
	case utils.NEGATION:
		return input.Input1.Neg().String(), nil
	case utils.SQUARE:
		power := decimal.NewFromInt(2)
		return input.Input1.Pow(power).String(), nil
	case utils.SQUAREROOT:
		return utils.Sqrt(input.Input1).String(), nil
	case utils.ABSOLUTE:
		return input.Input1.Abs().String(), nil
	case utils.CUBE:
		power := decimal.NewFromInt(3)
		return input.Input1.Pow(power).String(), nil
	case utils.CUBERT:
		floatInput, _ := input.Input1.Float64()
		return decimal.NewFromFloat(math.Cbrt(floatInput)).String(), nil
	default:
		return "", errors.ErrInvalidOperation
	}
}

func (c *calculatorService) validateAndConstructCalculationMultipleInput(inputs []string) (CalculationWithMultipleInput, errors.WrappedError) {
	if len(inputs) != 3 || !(utils.ContainString(utils.OPERATIONS_WITH_TWO_INPUTS, inputs[1])) {
		return CalculationWithMultipleInput{}, errors.ErrInvalidOperation
	}

	num1, err1 := decimal.NewFromString(inputs[0])
	num2, err2 := decimal.NewFromString(inputs[2])

	if !(err1 == nil && err2 == nil) {
		return CalculationWithMultipleInput{}, errors.ErrInvalidInputToBeOperated
	}
	return CalculationWithMultipleInput{Input1: num1, Input2: num2, Operation: inputs[1]}, nil
}

func (c *calculatorService) doCalculationWithMultipleInput(input CalculationWithMultipleInput) (string, errors.WrappedError) {
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
		return "", errors.ErrInvalidOperation
	}
}
