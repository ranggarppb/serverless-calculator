package calculation

import (
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/utils"
)

type CalculationInput struct {
	Input string `json:"input"`
}

func (i CalculationInput) ParseCalculationInput(history CalculationHistory) (ICalculationInput, errors.WrappedError) {
	if len(i.Input) == 0 {
		return nil, errors.ErrInvalidInput
	}

	inputs := strings.Split(i.Input, " ")

	if len(inputs) < 2 {
		return nil, errors.ErrInvalidOperation
	}

	switch {
	case utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, inputs[1]):
		filteredInputs := utils.Filter(append([]string{inputs[0]}, inputs[2:]...), func(s string) bool { return len(s) != 0 })
		return calculationWithOneInput{Operation: inputs[1], Inputs: filteredInputs}, nil
	case inputs[1] == utils.REPEAT:
		return calculationRepeatInput{Inputs: inputs[1:], Initial: inputs[0], History: history}, nil
	default:
		return calculationWithMultipleInput{Inputs: inputs}, nil
	}
}
