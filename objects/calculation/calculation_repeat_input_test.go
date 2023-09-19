package calculation_test

import (
	"testing"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculationRepeatInputTestSuite struct {
	suite.Suite
}

func TestSuiteCalculationRepeatInput(t *testing.T) {
	s := calculationRepeatInputTestSuite{}
	suite.Run(t, &s)
}

func (s *calculationRepeatInputTestSuite) TestGetInput() {
	history := calculation.CalculationHistory{
		History: []calculation.CalculationResult{
			{Input: "1 add 2", Result: "3"},
			{Input: "sqr 2", Result: "4"},
		},
	}
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedRes string
	}{
		{
			Desc:        "success-get-input-from-history-of-multiple-input-calculation",
			Input:       calculation.NewCalculationRepeatInput([]string{"repeat", "2"}, "2", history),
			ExpectedRes: "2 add 2",
		},
		{
			Desc:        "success-get-input-from-history-of-one-input-calculation",
			Input:       calculation.NewCalculationRepeatInput([]string{"repeat", "1"}, "3", history),
			ExpectedRes: " sqr 3",
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			res := tC.Input.GetInput()

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}

func (s *calculationRepeatInputTestSuite) TestValidate() {
	history := calculation.CalculationHistory{
		History: []calculation.CalculationResult{
			{Input: "1 add 2", Result: "3"},
			{Input: "sqr 2", Result: "4"},
		},
	}
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedErr errors.WrappedError
	}{
		{
			Desc:        "success-validate",
			Input:       calculation.NewCalculationRepeatInput([]string{"repeat", "2"}, "2", history),
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-validate-inputs-length-more-than-two",
			Input:       calculation.NewCalculationRepeatInput([]string{"repeat", "2", "random"}, "2", history),
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-repeat-input-more-than-history",
			Input:       calculation.NewCalculationRepeatInput([]string{"repeat", "10"}, "2", history),
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

func (s *calculationRepeatInputTestSuite) TestCalculate() {
	history := calculation.CalculationHistory{
		History: []calculation.CalculationResult{
			{Input: "1 add 2 multiply 3", Result: "7"},
			{Input: "sqr 2", Result: "4"},
		},
	}
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedRes string
	}{
		{
			Desc:        "success-calculate-with-multiple-input-calculation-history",
			Input:       calculation.NewCalculationRepeatInput([]string{"repeat", "2"}, "2", history),
			ExpectedRes: "8",
		},
		{
			Desc:        "success-calculate-with-multiple-one-calculation-history",
			Input:       calculation.NewCalculationRepeatInput([]string{"repeat", "1"}, "4", history),
			ExpectedRes: "16",
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			res := tC.Input.Calculate()

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}
