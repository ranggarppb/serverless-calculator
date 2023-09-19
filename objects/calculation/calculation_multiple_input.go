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
	if len(i.Inputs)%2 != 1 {
		return errors.ErrInvalidOperation
	}
	for idx, i := range i.Inputs {
		if idx%2 == 1 && !utils.ContainString(utils.OPERATIONS_WITH_MULTIPLE_INPUTS, i) {
			return errors.ErrInvalidOperation
		} else if _, err := decimal.NewFromString(i); idx%2 == 0 && err != nil {
			return errors.ErrInvalidInputToBeOperated
		}
	}

	return nil
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
