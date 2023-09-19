package calculator_test

import (
	"context"
	"testing"

	"github.com/ranggarppb/serverless-calculator/calculator"
	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/mocks"
	"github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculatorServiceTestSuite struct {
	suite.Suite
}

func TestSuiteCalculatorService(t *testing.T) {
	s := calculatorServiceTestSuite{}
	suite.Run(t, &s)
}

func (s *calculatorServiceTestSuite) TestCalculate() {
	testCases := []struct {
		Desc        string
		Ctx         context.Context
		Input       calculation.ICalculationInput
		ExpectedRes calculation.CalculationResult
		ExpectedErr errors.WrappedError
	}{
		{
			Desc: "success-calculate",
			Ctx:  context.TODO(),
			Input: func() calculation.ICalculationInput {
				input := new(mocks.ICalculationInput)
				input.On("Validate").Return(nil).
					On("Calculate").Return("4").
					On("GetInput").Return("1 add 3")
				return input
			}(),
			ExpectedRes: calculation.CalculationResult{Input: "1 add 3", Result: "4"},
			ExpectedErr: nil,
		},
		{
			Desc: "failed-calculate-failed-validate",
			Ctx:  context.TODO(),
			Input: func() calculation.ICalculationInput {
				input := new(mocks.ICalculationInput)
				input.On("Validate").Return(errors.ErrInvalidInputToBeOperated).
					On("Calculate").Return("").
					On("GetInput").Return("1 add random")
				return input
			}(),
			ExpectedRes: calculation.CalculationResult{},
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorService := calculator.NewCalculatorService()

			res, err := calculatorService.Calculate(tC.Ctx, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}
