package calculation_test

import (
	"testing"

	"github.com/ranggarppb/serverless-calculator/errors"
	calculator "github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculationTwoInputTestSuite struct {
	suite.Suite
}

func TestSuiteCalculationTwoInput(t *testing.T) {
	s := calculationTwoInputTestSuite{}
	suite.Run(t, &s)
}

func (s *calculationTwoInputTestSuite) TestGetInput() {
	testCases := []struct {
		Desc        string
		Input       calculator.ICalculationInput
		ExpectedRes string
	}{
		{
			Desc:        "success-get-input",
			Input:       calculator.NewCalculationWithTwoInput("2", "3", "add"),
			ExpectedRes: "2 add 3",
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			res := tC.Input.GetInput()

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}

func (s *calculationTwoInputTestSuite) TestValidate() {
	testCases := []struct {
		Desc        string
		Input       calculator.ICalculationInput
		ExpectedErr error
	}{
		{
			Desc:        "success-validate-calculation-with-two-input",
			Input:       calculator.NewCalculationWithTwoInput("2", "3", "add"),
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-validate-with-not-listed-operation",
			Input:       calculator.NewCalculationWithTwoInput("2", "3", "random"),
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-with-invalid-input-to-be-operated",
			Input:       calculator.NewCalculationWithTwoInput("random", "3", "add"),
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			err := tC.Input.Validate()

			require.Equal(t, tC.ExpectedErr, err)
		})
	}
}

func (s *calculationTwoInputTestSuite) TestCalculate() {
	input1 := "2"
	input1Dec, _ := decimal.NewFromString(input1)
	input2 := "3"
	input2Dec, _ := decimal.NewFromString(input2)
	testCases := []struct {
		Desc        string
		Input       calculator.ICalculationInput
		ExpectedRes string
	}{
		{
			Desc:        "success-do-addition",
			Input:       calculator.NewCalculationWithTwoInput(input1, input2, "add"),
			ExpectedRes: input1Dec.Add(input2Dec).String(),
		},
		{
			Desc:        "success-do-negation",
			Input:       calculator.NewCalculationWithTwoInput(input1, input2, "subtract"),
			ExpectedRes: input1Dec.Sub(input2Dec).String(),
		},
		{
			Desc:        "success-do-multiplicaiton",
			Input:       calculator.NewCalculationWithTwoInput(input1, input2, "multiply"),
			ExpectedRes: input1Dec.Mul(input2Dec).String(),
		},
		{
			Desc:        "success-do-division",
			Input:       calculator.NewCalculationWithTwoInput(input1, input2, "divide"),
			ExpectedRes: input1Dec.Div(input2Dec).String(),
		},
		{
			Desc:        "return-empty-for-invalid-operation",
			Input:       calculator.NewCalculationWithTwoInput(input1, input2, "random"),
			ExpectedRes: "",
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			res := tC.Input.Calculate()

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}
