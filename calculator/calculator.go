package calculator

import (
	"context"
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	cl "github.com/ranggarppb/serverless-calculator/objects/calculation"
)

type calculatorService struct{}

func NewCalculatorService() *calculatorService {
	return &calculatorService{}
}

func (c *calculatorService) GetCalculationHistory(ctx context.Context) cl.CalculationHistory {

	return cl.CalculationHistory{}
}

func (c *calculatorService) Calculate(ctx context.Context, input cl.ICalculationInput) (cl.CalculationResult, errors.WrappedError) {
	if err := input.Validate(); err != nil {
		return cl.CalculationResult{}, err
	}

	return cl.CalculationResult{Input: strings.TrimPrefix(input.GetInput(), " "), Result: input.Calculate()}, nil
}
