package calculator_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/ranggarppb/serverless-calculator/calculator"
	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/types/structs"
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
	input1Float64, _ := input1Dec.Float64()
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
			ExpectedRes: decimal.NewFromFloat(math.Sqrt(input1Float64)).String(),
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
	input2 := "3"
	testCases := []struct {
		Desc        string
		Input       []string
		ExpectedRes structs.CalculationWithMultipleInput
		ExpectedErr error
	}{
		{
			Desc:        "success-validate-and-construct-calculation-with-multiple-input",
			Input:       []string{input1, "add", input2},
			ExpectedRes: structs.CalculationWithMultipleInput{Inputs: []string{input1, "add", input2}},
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

func (s *calculatorTestSuite) TestDoCalculationWithTwoInput() {
	input1 := "2"
	input1Dec, _ := decimal.NewFromString(input1)
	input2 := "3"
	input2Dec, _ := decimal.NewFromString(input2)
	testCases := []struct {
		Desc        string
		Input       structs.CalculationWithTwoInput
		ExpectedRes string
		ExpectedErr error
	}{
		{
			Desc:        "success-do-addition",
			Input:       structs.CalculationWithTwoInput{Input1: input1Dec, Input2: input2Dec, Operation: "add"},
			ExpectedRes: input1Dec.Add(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-negation",
			Input:       structs.CalculationWithTwoInput{Input1: input1Dec, Input2: input2Dec, Operation: "subtract"},
			ExpectedRes: input1Dec.Sub(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-multiplicaiton",
			Input:       structs.CalculationWithTwoInput{Input1: input1Dec, Input2: input2Dec, Operation: "multiply"},
			ExpectedRes: input1Dec.Mul(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-do-division",
			Input:       structs.CalculationWithTwoInput{Input1: input1Dec, Input2: input2Dec, Operation: "divide"},
			ExpectedRes: input1Dec.Div(input2Dec).String(),
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-with-invalid-operation",
			Input:       structs.CalculationWithTwoInput{Input1: input1Dec, Input2: input2Dec, Operation: "random"},
			ExpectedRes: "",
			ExpectedErr: errors.ErrInvalidOperation,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetDoCalculationWithTwoInput(calculatorSvc, tC.Input)

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
			ExpectedRes: structs.CalculationWithMultipleInput{Inputs: []string{input1, "add", input2}},
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parsing-with-normal-input-for-calculation-with-multiple-input-2",
			Input:       fmt.Sprintf("%v add %v multiply %v", input1, input2, input2),
			ExpectedRes: structs.CalculationWithMultipleInput{Inputs: []string{input1, "add", input2, "multiply", input2}},
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

func (s *calculatorTestSuite) TestChangeToPostfixOperation() {
	testCases := []struct {
		Desc        string
		Inputs      []string
		ExpectedRes []string
	}{
		{
			Desc:        "success-change-two-inputs-add-operation-to-postfix",
			Inputs:      []string{"1", "add", "2"},
			ExpectedRes: []string{"1", "2", "add"},
		},
		{
			Desc:        "success-change-two-inputs-subtract-operation-to-postfix",
			Inputs:      []string{"1", "subtract", "2"},
			ExpectedRes: []string{"1", "2", "subtract"},
		},
		{
			Desc:        "success-change-two-inputs-multiply-operation-to-postfix",
			Inputs:      []string{"1", "multiply", "2"},
			ExpectedRes: []string{"1", "2", "multiply"},
		},
		{
			Desc:        "success-change-two-inputs-divide-operation-to-postfix",
			Inputs:      []string{"1", "divide", "2"},
			ExpectedRes: []string{"1", "2", "divide"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-equal-priority-operation-1-to-postfix",
			Inputs:      []string{"1", "add", "3", "subtract", "5"},
			ExpectedRes: []string{"1", "3", "add", "5", "subtract"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-equal-priority-operation-2-to-postfix",
			Inputs:      []string{"1", "multiply", "3", "divide", "5"},
			ExpectedRes: []string{"1", "3", "multiply", "5", "divide"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-1-to-postfix",
			Inputs:      []string{"1", "add", "3", "multiply", "5"},
			ExpectedRes: []string{"1", "3", "5", "multiply", "add"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-2-to-postfix",
			Inputs:      []string{"1", "multiply", "3", "add", "5"},
			ExpectedRes: []string{"1", "3", "multiply", "5", "add"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-3-to-postfix",
			Inputs:      []string{"1", "add", "3", "multiply", "5", "subtract", "2"},
			ExpectedRes: []string{"1", "3", "5", "multiply", "add", "2", "subtract"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-4-to-postfix",
			Inputs:      []string{"1", "multiply", "3", "add", "5", "divide", "2"},
			ExpectedRes: []string{"1", "3", "multiply", "5", "2", "divide", "add"},
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res := calculator.GetChangeToPostfixOperation(calculatorSvc, tC.Inputs)

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}

func (s *calculatorTestSuite) TestCalculatePostfixOperation() {
	testCases := []struct {
		Desc        string
		Inputs      []string
		ExpectedRes string
	}{
		{
			Desc:        "success-calculate-two-inputs-add-postfix-operation",
			Inputs:      []string{"1", "2", "add"},
			ExpectedRes: "3",
		},
		{
			Desc:        "success-calculate-two-inputs-subtract-postfix-operation",
			Inputs:      []string{"1", "2", "subtract"},
			ExpectedRes: "-1",
		},
		{
			Desc:        "success-calculate-two-inputs-multiply-postfix-operation",
			Inputs:      []string{"1", "2", "multiply"},
			ExpectedRes: "2",
		},
		{
			Desc:        "success-calculate-two-inputs-divide-postfix-operation",
			Inputs:      []string{"1", "2", "divide"},
			ExpectedRes: "0.5",
		},
		{
			Desc:        "success-calculate-multiple-inputs-with-equal-priority-postfix-operation-1",
			Inputs:      []string{"1", "3", "add", "5", "subtract"},
			ExpectedRes: "-1",
		},
		{
			Desc:        "success-calculate-multiple-inputs-with-equal-priority-postfix-operation-2",
			Inputs:      []string{"15", "3", "multiply", "5", "divide"},
			ExpectedRes: "9",
		},
		{
			Desc:        "success-calculate-multiple-inputs-with-non-equal-priority-postfix-operation-1",
			Inputs:      []string{"1", "3", "5", "multiply", "add"},
			ExpectedRes: "16",
		},
		{
			Desc:        "success-calculate-multiple-inputs-with-non-equal-priority-postfix-operation-2",
			Inputs:      []string{"1", "3", "multiply", "5", "add"},
			ExpectedRes: "8",
		},
		{
			Desc:        "success-calculate-multiple-inputs-with-non-equal-priority-postfix-operation-3",
			Inputs:      []string{"1", "3", "5", "multiply", "add", "2", "subtract"},
			ExpectedRes: "14",
		},
		{
			Desc:        "success-calculate-multiple-inputs-with-non-equal-priority-postfix-operation-4",
			Inputs:      []string{"1", "3", "multiply", "6", "2", "divide", "add"},
			ExpectedRes: "6",
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, _ := calculator.GetCalculatePostfixOperation(calculatorSvc, tC.Inputs)

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}
