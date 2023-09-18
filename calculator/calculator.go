package calculator

import (
	"context"
	"math"
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	cl "github.com/ranggarppb/serverless-calculator/types/calculator"
	"github.com/ranggarppb/serverless-calculator/utils"
	"github.com/shopspring/decimal"
)

type calculatorService struct{}

func NewCalculatorService() *calculatorService {
	return &calculatorService{}
}

func (c *calculatorService) GetCalculationHistory(ctx context.Context) cl.CalculationHistory {

	return cl.CalculationHistory{}
}

func (c *calculatorService) Calculate(ctx context.Context, input string) (cl.CalculationResult, errors.WrappedError) {

	parsedInput, err := c.parseInput(input)

	if err != nil {
		return cl.CalculationResult{}, err
	}

	switch inputType := parsedInput.(type) {
	case cl.CalculationWithOneInput:
		result, err := c.doCalculationWithOneInput(inputType)
		if err != nil {
			return cl.CalculationResult{}, err
		}
		return cl.CalculationResult{Input: input, Result: result}, nil
	case cl.CalculationWithMultipleInput:
		result, err := c.doCalculationWithMultipleInput(inputType)
		if err != nil {
			return cl.CalculationResult{}, err
		}
		return cl.CalculationResult{Input: input, Result: result}, nil
	default:
		return cl.CalculationResult{}, errors.ErrInvalidOperation
	}
}

func (c *calculatorService) parseInput(input string) (interface{}, errors.WrappedError) {
	inputs := strings.Split(input, " ")

	calculationType, err := c.getCalculationType(inputs)

	if err != nil {
		return nil, err
	}

	switch calculationType.(type) {
	case cl.CalculationWithOneInput:
		res, err := c.validateAndConstructCalculationOneInput(inputs)

		if err != nil {
			return cl.CalculationWithOneInput{}, err
		}

		return res, nil
	case cl.CalculationWithMultipleInput:
		res, err := c.validateAndConstructCalculationMultipleInput(inputs)

		if err != nil {
			return cl.CalculationWithMultipleInput{}, err
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
		return cl.CalculationWithOneInput{}, nil
	} else {
		return cl.CalculationWithMultipleInput{}, nil
	}
}

func (c *calculatorService) validateAndConstructCalculationOneInput(inputs []string) (cl.CalculationWithOneInput, errors.WrappedError) {
	if len(inputs) != 2 || !(utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, inputs[0])) {
		return cl.CalculationWithOneInput{}, errors.ErrInvalidOperation
	}

	num, err := decimal.NewFromString(inputs[1])

	if err != nil {
		return cl.CalculationWithOneInput{}, errors.ErrInvalidInputToBeOperated
	}
	if inputs[0] == utils.SQUAREROOT && num.LessThan(decimal.Zero) {
		return cl.CalculationWithOneInput{}, errors.ErrInvalidInputToBeOperated
	}
	return cl.CalculationWithOneInput{Input1: num, Operation: inputs[0]}, nil
}

func (c *calculatorService) doCalculationWithOneInput(input cl.CalculationWithOneInput) (string, errors.WrappedError) {
	switch input.Operation {
	case utils.NEGATION:
		return input.Input1.Neg().String(), nil
	case utils.SQUARE:
		power := decimal.NewFromInt(2)
		return input.Input1.Pow(power).String(), nil
	case utils.SQUAREROOT:
		floatInput, _ := input.Input1.Float64()
		return decimal.NewFromFloat(math.Sqrt(floatInput)).String(), nil
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

func (c *calculatorService) validateAndConstructCalculationMultipleInput(inputs []string) (cl.CalculationWithMultipleInput, errors.WrappedError) {
	if len(inputs)%2 != 1 {
		return cl.CalculationWithMultipleInput{}, errors.ErrInvalidOperation
	}
	for idx, i := range inputs {
		if idx%2 == 1 && !utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, i) {
			return cl.CalculationWithMultipleInput{}, errors.ErrInvalidOperation
		} else if _, err := decimal.NewFromString(i); idx%2 == 0 && err != nil {
			return cl.CalculationWithMultipleInput{}, errors.ErrInvalidInputToBeOperated
		}
	}
	return cl.CalculationWithMultipleInput{Inputs: inputs}, nil
}

func (c *calculatorService) doCalculationWithTwoInput(input cl.CalculationWithTwoInput) (string, errors.WrappedError) {
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

func (c *calculatorService) doCalculationWithMultipleInput(input cl.CalculationWithMultipleInput) (string, errors.WrappedError) {
	postfixOperation := c.changeToPostfixOperation(input.Inputs)

	return c.calculatePostfixOperation(postfixOperation)
}

func (c *calculatorService) changeToPostfixOperation(inputs []string) []string {
	operationStacks := []string{}
	res := []string{}

	for _, i := range inputs {
		if utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, i) {
			if len(operationStacks) == 0 {
				operationStacks = append(operationStacks, i)
			} else if c.isMorePriority(i, operationStacks[len(operationStacks)-1]) {
				operationStacks = append(operationStacks, i)
			} else {
				res = append(res, utils.Revert(operationStacks)...)
				operationStacks = []string{i}
			}
		} else {
			res = append(res, i)
		}
	}

	if len(operationStacks) > 0 {
		res = append(res, utils.Revert(operationStacks)...)
	}

	return res
}

func (c *calculatorService) calculatePostfixOperation(inputs []string) (string, errors.WrappedError) {

	resStack := []string{}
	for _, i := range inputs {
		if utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, i) {
			operand1, _ := decimal.NewFromString(resStack[len(resStack)-2])
			operand2, _ := decimal.NewFromString(resStack[len(resStack)-1])
			operation := cl.CalculationWithTwoInput{Input1: operand1, Input2: operand2, Operation: i}
			result, err := c.doCalculationWithTwoInput(operation)

			if err != nil {
				return "", err
			}

			resStack = resStack[:len(resStack)-2]
			resStack = append(resStack, result)
		} else {
			resStack = append(resStack, i)
		}
	}

	return resStack[0], nil
}

func (c *calculatorService) isMorePriority(operation1 string, operation2 string) bool {
	return utils.OPERATIONS_PRIORITY[operation1] > utils.OPERATIONS_PRIORITY[operation2]
}
