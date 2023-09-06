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

	return input, nil
}

func (c *calculatorService) parseInput(input string) ([]interface{}, types.WrappedError) {
	inputs := strings.Split(input, " ")

	if len(inputs) != 3 || !(utils.ContainString(utils.ALLOWED_OPERATIONS, inputs[1])) {
		return []interface{}{}, types.ErrInvalidOperation
	}

	num1, err := decimal.NewFromString(inputs[0])
	if err != nil {
		return []interface{}{}, types.ErrInvalidInputToBeOperated
	}
	num2, err := decimal.NewFromString(inputs[2])
	if err != nil {
		return []interface{}{}, types.ErrInvalidInputToBeOperated
	}

	res := []interface{}{num1, inputs[1], num2}

	return res, nil
}
