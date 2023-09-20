package calculation

import (
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/utils"
	"github.com/shopspring/decimal"
)

type calculationWithMultipleInput struct {
	Inputs []string
}

func NewCalculationWithMultipleInput(inputs []string) calculationWithMultipleInput {
	return calculationWithMultipleInput{Inputs: inputs}
}

func (i calculationWithMultipleInput) GetInput() string {
	return strings.Join(i.Inputs, " ")
}

func (i calculationWithMultipleInput) Validate() errors.WrappedError {
	for idx, element := range i.Inputs {
		_, err := decimal.NewFromString(element)

		switch {
		case err == nil:
			if (idx == 0) && !i.validOperation(i.Inputs[idx+1]) {
				return errors.ErrInvalidOperation
			}
			if (idx == len(i.Inputs)-1) && !i.validOperation(i.Inputs[idx-1]) {
				return errors.ErrInvalidOperation
			}
			if !(idx == 0 || idx == len(i.Inputs)-1) && !i.validOperation(i.Inputs[idx-1]) {
				return errors.ErrInvalidOperation
			}
			if !(idx == 0 || idx == len(i.Inputs)-1) && !i.validOperation(i.Inputs[idx+1]) {
				return errors.ErrInvalidOperation
			}
		case utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, element):
			if (idx == 0) || (idx == len(i.Inputs)-1) {
				return errors.ErrInvalidOperation
			}

			if !(idx == 0) || (idx == len(i.Inputs)-1) {
				_, err1 := decimal.NewFromString(i.Inputs[idx-1])
				_, err2 := decimal.NewFromString(i.Inputs[idx+1])
				if !(err1 == nil || err2 == nil) {
					return errors.ErrInvalidInputToBeOperated
				}
			}
		default:
			if idx == 0 {
				return errors.ErrInvalidOperation
			}
			if idx == len(i.Inputs)-1 {
				return errors.ErrInvalidInputToBeOperated
			}
			if !(idx == 0) || (idx == len(i.Inputs)-1) && i.validOperation(i.Inputs[idx-1]) {
				return errors.ErrInvalidInputToBeOperated
			}
			return errors.ErrInvalidOperation
		}
	}

	return nil
}

func (i calculationWithMultipleInput) validOperation(input string) bool {
	return utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, input)
}

func (i calculationWithMultipleInput) Calculate() string {
	postfixOperation := i.changeToPostfixOperation(i.Inputs)

	return i.calculatePostfixOperation(postfixOperation)
}

func (i calculationWithMultipleInput) changeToPostfixOperation(inputs []string) []string {
	operationStacks := []string{}
	res := []string{}

	for _, input := range inputs {
		if utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, input) {
			if len(operationStacks) == 0 {
				operationStacks = append(operationStacks, input)
			} else if i.isMorePriority(input, operationStacks[len(operationStacks)-1]) {
				operationStacks = append(operationStacks, input)
			} else if i.isSamePriority(input, operationStacks[len(operationStacks)-1]) {
				res = append(res, operationStacks[len(operationStacks)-1])
				operationStacks = operationStacks[:len(operationStacks)-1]
				operationStacks = append(operationStacks, input)
			} else {
				res = append(res, utils.Revert(operationStacks)...)
				operationStacks = []string{input}
			}
		} else {
			res = append(res, input)
		}
	}

	if len(operationStacks) > 0 {
		res = append(res, utils.Revert(operationStacks)...)
	}

	return res
}

func (c calculationWithMultipleInput) calculatePostfixOperation(inputs []string) string {

	resStack := []string{}
	for _, i := range inputs {
		if utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, i) {
			calculationWithTwoInput := calculationWithTwoInput{Input1: resStack[len(resStack)-2], Input2: resStack[len(resStack)-1], Operation: i}
			result := calculationWithTwoInput.Calculate()

			resStack = resStack[:len(resStack)-2]
			resStack = append(resStack, result)
		} else {
			resStack = append(resStack, i)
		}
	}

	return resStack[0]
}

func (i calculationWithMultipleInput) isMorePriority(operation1 string, operation2 string) bool {
	return utils.OPERATIONS_PRIORITY[operation1] > utils.OPERATIONS_PRIORITY[operation2]
}

func (i calculationWithMultipleInput) isSamePriority(operation1 string, operation2 string) bool {
	return utils.OPERATIONS_PRIORITY[operation1] == utils.OPERATIONS_PRIORITY[operation2]
}
