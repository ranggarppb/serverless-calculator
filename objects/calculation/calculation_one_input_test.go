package calculation_test

import (
	"math"
	"testing"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculationOneInputTestSuite struct {
	suite.Suite
}

func TestSuiteCalculationOneInput(t *testing.T) {
	s := calculationOneInputTestSuite{}
	suite.Run(t, &s)
}

func (s *calculationOneInputTestSuite) TestGetInput() {
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedRes string
	}{
		{
			Desc:        "success-get-input-1",
			Input:       calculation.NewCalculationWithOneInput("sqr", []string{"2"}),
			ExpectedRes: "sqr 2",
		},
		{
			Desc:        "success-get-input-2",
			Input:       calculation.NewCalculationWithOneInput("sqr", []string{"2", "3"}),
			ExpectedRes: "sqr 2 3",
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			res := tC.Input.GetInput()

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}

func (s *calculationOneInputTestSuite) TestValidateInput() {
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedErr error
	}{
		{
			Desc:        "success-validate-input",
			Input:       calculation.NewCalculationWithOneInput("sqr", []string{"2"}),
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-validate-input-length-more-than-1",
			Input:       calculation.NewCalculationWithOneInput("sqr", []string{"2", "3"}),
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-with-not-listed-operation",
			Input:       calculation.NewCalculationWithOneInput("random", []string{"2"}),
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-with-invalid-input-to-be-operated",
			Input:       calculation.NewCalculationWithOneInput("sqr", []string{"random"}),
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
		{
			Desc:        "failed-validate-when-input-is-negative-and-operation-is-sqrt",
			Input:       calculation.NewCalculationWithOneInput("sqrt", []string{"-2"}),
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

func (s *calculationOneInputTestSuite) TestCalculate() {
	input := "2"
	inputDec, _ := decimal.NewFromString(input)
	inputFloat64, _ := inputDec.Float64()
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedRes string
	}{
		{
			Desc:        "success-do-square",
			Input:       calculation.NewCalculationWithOneInput("sqr", []string{input}),
			ExpectedRes: inputDec.Pow(decimal.NewFromInt(2)).String(),
		},
		{

			Desc:        "success-do-negation",
			Input:       calculation.NewCalculationWithOneInput("neg", []string{input}),
			ExpectedRes: inputDec.Neg().String(),
		},
		{
			Desc:        "success-do-squareroot",
			Input:       calculation.NewCalculationWithOneInput("sqrt", []string{input}),
			ExpectedRes: decimal.NewFromFloat(math.Sqrt(inputFloat64)).String(),
		},
		{
			Desc:        "success-do-abs",
			Input:       calculation.NewCalculationWithOneInput("abs", []string{input}),
			ExpectedRes: inputDec.Abs().String(),
		},
		{
			Desc:        "success-do-cube",
			Input:       calculation.NewCalculationWithOneInput("cube", []string{input}),
			ExpectedRes: inputDec.Pow(decimal.NewFromInt(3)).String(),
		},
		{
			Desc:        "success-do-cubert",
			Input:       calculation.NewCalculationWithOneInput("cubert", []string{"8"}),
			ExpectedRes: "2",
		},
		{
			Desc:        "return-empty-with-invalid-operation",
			Input:       calculation.NewCalculationWithOneInput("random", []string{"2"}),
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
