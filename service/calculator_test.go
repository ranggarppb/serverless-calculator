package calculator_test

import (
	"fmt"
	"testing"

	calculator "github.com/ranggarppb/serverless-calculator/service"
	"github.com/ranggarppb/serverless-calculator/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculatorTestSuite struct {
	suite.Suite
}

func TestSuiteCalculator(t *testing.T) {
	s := calculatorTestSuite{}
	suite.Run(t, &s)
}

func (s *calculatorTestSuite) TestParsingInput() {
	input1 := "1"
	input2 := "2"
	input1Dec, _ := decimal.NewFromString(input1)
	input2Dec, _ := decimal.NewFromString(input2)
	testCases := []struct {
		Desc        string
		Input       string
		ExpectedRes types.CalculationWithTwoInput
		ExpectedErr error
	}{
		{
			Desc:        "success-parsing-with-normal-input",
			Input:       fmt.Sprintf("%v add %v", input1, input2),
			ExpectedRes: types.CalculationWithTwoInput{Input1: input1Dec, Input2: input2Dec, Operation: "add"},
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-parsing-with-inputs-not-separated-by-space",
			Input:       fmt.Sprintf("%vadd %v", input1, input2),
			ExpectedRes: types.CalculationWithTwoInput{},
			ExpectedErr: types.ErrInvalidOperation,
		},
		{
			Desc:        "failed-parsing-when-containing-not-allowed-operation",
			Input:       "1 random 2",
			ExpectedRes: types.CalculationWithTwoInput{},
			ExpectedErr: types.ErrInvalidOperation,
		},
		{
			Desc:        "failed-parsing-when-containing-input-cannot-be-operated",
			Input:       "add add 2",
			ExpectedRes: types.CalculationWithTwoInput{},
			ExpectedErr: types.ErrInvalidInputToBeOperated,
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetParseInput(calculatorSvc, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}
