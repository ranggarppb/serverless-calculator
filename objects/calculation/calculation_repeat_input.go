package calculation

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/utils"
)

type calculationRepeatInput struct {
	Inputs  []string
	Initial string
	History CalculationHistory
}

func NewCalculationRepeatInput(inputs []string, initial string, history CalculationHistory) calculationRepeatInput {
	return calculationRepeatInput{Inputs: inputs, Initial: initial, History: history}
}

func (i calculationRepeatInput) GetInput() string {
	repeatInput, _ := strconv.Atoi(i.Inputs[1])
	operationToBeRepeated := strings.Split(i.History.History[len(i.History.History)-repeatInput].Input, " ")

	if utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, operationToBeRepeated[0]) {
		return fmt.Sprintf(" %s %s", operationToBeRepeated[0], i.Initial)
	} else {
		operationToBeRepeated[0] = i.Initial
		return strings.Join(operationToBeRepeated, " ")
	}

}

func (i calculationRepeatInput) Validate() errors.WrappedError {
	if len(i.Inputs) > 2 {
		return errors.ErrInvalidOperation
	}
	repeatInput, err := strconv.Atoi(i.Inputs[1])

	if err != nil || repeatInput > len(i.History.History) {
		return errors.ErrInvalidInputToBeOperated
	}

	return nil
}

func (i calculationRepeatInput) Calculate() string {
	repeatInput := CalculationInput{Input: i.GetInput()}
	parsedRepeatInput, _ := repeatInput.ParseCalculationInput(CalculationHistory{})

	return parsedRepeatInput.Calculate()
}
