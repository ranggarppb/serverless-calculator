package calculator_test

import (
	"fmt"
	"testing"

	"github.com/ranggarppb/serverless-calculator/calculator"
	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/types/structs"
	"github.com/ranggarppb/serverless-calculator/utils"
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

func (s *calculatorTestSuite) TestGetCalculationType() {
	testCases := []struct {
		Desc        string
		Input       []string
		ExpectedRes interface{}
		ExpectedErr error
	}{
		{
			Desc:        "success-get-type-with-normal-input-for-calculation-with-one-input",
			Input:       []string{"sqr", "2"},
			ExpectedRes: structs.CalculationWithOneInput{},
			ExpectedErr: nil,
		},
		{
			Desc:        "success-get-type-with-normal-input-for-calculation-with-one-input",
			Input:       []string{"1", "add", "2"},
			ExpectedRes: structs.CalculationWithMultipleInput{},
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-get-calculation-type-with-empty-inputs",
			Input:       []string{},
			ExpectedRes: nil,
			ExpectedErr: errors.ErrInvalidInput,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetGetCalculationType(calculatorSvc, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}

func (s *calculatorTestSuite) TestValidateAndConstructCalculationOneInput() {
	input := "2"
	input1Dec, _ := decimal.NewFromString(input)
	testCases := []struct {
		Desc        string
		Input       []string
		ExpectedRes structs.CalculationWithOneInput
		ExpectedErr error
	}{
		{
			Desc:        "success-validate-and-construct-calculation-with-one-input",
			Input:       []string{"sqr", input},
			ExpectedRes: structs.CalculationWithOneInput{Input1: input1Dec, Operation: "sqr"},
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-validate-and-construct-with-input-length-not-equal-to-two",
			Input:       []string{"sqr2"},
			ExpectedRes: structs.CalculationWithOneInput{},
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-and-construct-with-not-listed-operation",
			Input:       []string{"random", input},
			ExpectedRes: structs.CalculationWithOneInput{},
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-and-construct-with-invalid-input-to-be-operated",
			Input:       []string{"sqr", "random"},
			ExpectedRes: structs.CalculationWithOneInput{},
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
		{
			Desc:        "failed-validate-and-construct-with-when-input-is-negative-and-operation-is-sqrt",
			Input:       []string{"sqrt", "-2"},
			ExpectedRes: structs.CalculationWithOneInput{},
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetValidateAndConstructCalculationOneInput(calculatorSvc, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}

func (s *calculatorTestSuite) TestDoCalculationWithOneInput() {
	input := "2"
	input1Dec, _ := decimal.NewFromString(input)
	testCases := []struct {
		Desc        string
		Input       structs.CalculationWithOneInput
		ExpectedRes string
		ExpectedErr error
	}{
		{
			Desc:        "success-do-square",
			Input:       structs.CalculationWithOneInput{Input1: input1Dec, Operation: "sqr"},
			ExpectedRes: input1Dec.Pow(decimal.NewFromInt(2)).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-negation",
			Input:       structs.CalculationWithOneInput{Input1: input1Dec, Operation: "neg"},
			ExpectedRes: input1Dec.Neg().String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-squareroot",
			Input:       structs.CalculationWithOneInput{Input1: input1Dec, Operation: "sqrt"},
			ExpectedRes: utils.Sqrt(input1Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-abs",
			Input:       structs.CalculationWithOneInput{Input1: input1Dec, Operation: "abs"},
			ExpectedRes: input1Dec.Abs().String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-cube",
			Input:       structs.CalculationWithOneInput{Input1: input1Dec, Operation: "cube"},
			ExpectedRes: input1Dec.Pow(decimal.NewFromInt(3)).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-cubert",
			Input:       structs.CalculationWithOneInput{Input1: decimal.NewFromInt(8), Operation: "cubert"},
			ExpectedRes: "2",
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-with-invalid-operation",
			Input:       structs.CalculationWithOneInput{Input1: input1Dec, Operation: "random"},
			ExpectedRes: "",
			ExpectedErr: errors.ErrInvalidOperation,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetDoCalculationWithOneInput(calculatorSvc, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}

func (s *calculatorTestSuite) TestValidateAndConstructCalculationMultipleInput() {
	input1 := "2"
	input1Dec, _ := decimal.NewFromString(input1)
	input2 := "3"
	input2Dec, _ := decimal.NewFromString(input2)
	testCases := []struct {
		Desc        string
		Input       []string
		ExpectedRes structs.CalculationWithMultipleInput
		ExpectedErr error
	}{
		{
			Desc:        "success-validate-and-construct-calculation-with-multiple-input",
			Input:       []string{input1, "add", input2},
			ExpectedRes: structs.CalculationWithMultipleInput{Input1: input1Dec, Input2: input2Dec, Operation: "add"},
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-validate-and-construct-with-input-length-not-equal-to-two",
			Input:       []string{"2", "add"},
			ExpectedRes: structs.CalculationWithMultipleInput{},
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-and-construct-with-not-listed-operation",
			Input:       []string{input1, "random", input2},
			ExpectedRes: structs.CalculationWithMultipleInput{},
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-and-construct-with-invalid-input-to-be-operated",
			Input:       []string{input1, "add", "random"},
			ExpectedRes: structs.CalculationWithMultipleInput{},
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetValidateAndConstructCalculationMultipleInput(calculatorSvc, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}

func (s *calculatorTestSuite) TestDoCalculationWithMultipleInput() {
	input1 := "2"
	input1Dec, _ := decimal.NewFromString(input1)
	input2 := "3"
	input2Dec, _ := decimal.NewFromString(input2)
	testCases := []struct {
		Desc        string
		Input       structs.CalculationWithMultipleInput
		ExpectedRes string
		ExpectedErr error
	}{
		{
			Desc:        "success-do-addition",
			Input:       structs.CalculationWithMultipleInput{Input1: input1Dec, Input2: input2Dec, Operation: "add"},
			ExpectedRes: input1Dec.Add(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-negation",
			Input:       structs.CalculationWithMultipleInput{Input1: input1Dec, Input2: input2Dec, Operation: "substract"},
			ExpectedRes: input1Dec.Sub(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-multiplicaiton",
			Input:       structs.CalculationWithMultipleInput{Input1: input1Dec, Input2: input2Dec, Operation: "multiply"},
			ExpectedRes: input1Dec.Mul(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-division",
			Input:       structs.CalculationWithMultipleInput{Input1: input1Dec, Input2: input2Dec, Operation: "divide"},
			ExpectedRes: input1Dec.Div(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-with-invalid-operation",
			Input:       structs.CalculationWithMultipleInput{Input1: input1Dec, Input2: input2Dec, Operation: "random"},
			ExpectedRes: "",
			ExpectedErr: errors.ErrInvalidOperation,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetDoCalculationWithMultipleInput(calculatorSvc, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}

func (s *calculatorTestSuite) TestParsingInput() {
	input1 := "1"
	input2 := "2"
	input1Dec, _ := decimal.NewFromString(input1)
	input2Dec, _ := decimal.NewFromString(input2)
	testCases := []struct {
		Desc        string
		Input       string
		ExpectedRes interface{}
		ExpectedErr error
	}{
		{
			Desc:        "success-parsing-with-normal-input-for-calculation-with-one-input",
			Input:       fmt.Sprintf("sqr %v", input1),
			ExpectedRes: structs.CalculationWithOneInput{Input1: input1Dec, Operation: "sqr"},
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parsing-with-normal-input-for-calculation-with-multiple-input",
			Input:       fmt.Sprintf("%v add %v", input1, input2),
			ExpectedRes: structs.CalculationWithMultipleInput{Input1: input1Dec, Input2: input2Dec, Operation: "add"},
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-calculation-with-one-input-invalid-input-to-be-operated",
			Input:       "sqr random",
			ExpectedRes: "",
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
		{
			Desc:        "failed-calculation-with-multiple-input-invalid-input-to-be-operated",
			Input:       "1 add random",
			ExpectedRes: "",
			ExpectedErr: errors.ErrInvalidInputToBeOperated,
		},
		{
			Desc:        "failed-calculation-with-one-input-invalid-invalid-operation",
			Input:       "random 1",
			ExpectedRes: "",
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-calculation-with-multiple-input-invalid-operation",
			Input:       "1 random 3",
			ExpectedRes: "",
			ExpectedErr: errors.ErrInvalidOperation,
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
