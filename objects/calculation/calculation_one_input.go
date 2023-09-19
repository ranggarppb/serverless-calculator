package calculation

import (
	"fmt"
	"math"
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/utils"
	"github.com/shopspring/decimal"
)

type calculationWithOneInput struct {
	Operation string
	Inputs    []string
}

func NewCalculationWithOneInput(operation string, inputs []string) calculationWithOneInput {
	return calculationWithOneInput{Operation: operation, Inputs: inputs}
}

func (i calculationWithOneInput) GetInput() string {
	return fmt.Sprintf("%s %s", i.Operation, strings.Join(i.Inputs, " "))
}

func (i calculationWithOneInput) Validate() errors.WrappedError {
	if len(i.Inputs) != 1 {
		return errors.ErrInvalidOperation
	}

	if !utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, i.Operation) {
		return errors.ErrInvalidOperation
	}

	num, err := decimal.NewFromString(i.Inputs[0])

	if err != nil {
		return errors.ErrInvalidInputToBeOperated
	}

	if i.Operation == utils.SQUAREROOT && num.LessThan(decimal.Zero) {
		return errors.ErrInvalidInputToBeOperated
	}

	return nil
}

func (i calculationWithOneInput) Calculate() string {

	input, _ := decimal.NewFromString(i.Inputs[0])

	switch i.Operation {
	case utils.NEGATION:
		return input.Neg().String()
	case utils.SQUARE:
		power := decimal.NewFromInt(2)
		return input.Pow(power).String()
	case utils.SQUAREROOT:
		floatInput, _ := input.Float64()
		return decimal.NewFromFloat(math.Sqrt(floatInput)).String()
	case utils.ABSOLUTE:
		return input.Abs().String()
	case utils.CUBE:
		power := decimal.NewFromInt(3)
		return input.Pow(power).String()
	case utils.CUBERT:
		floatInput, _ := input.Float64()
		return decimal.NewFromFloat(math.Cbrt(floatInput)).String()
	default:
		return ""
	}
}
