package calculation_test

import (
	"testing"

	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculationMultipleInputTestSuite struct {
	suite.Suite
}

func TestSuiteCalculationMultipleInput(t *testing.T) {
	s := calculationMultipleInputTestSuite{}
	suite.Run(t, &s)
}

func (s *calculationMultipleInputTestSuite) TestGetInput() {
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedRes string
	}{
		{
			Desc:        "success-get-input",
			Input:       calculation.NewCalculationWithMultipleInput([]string{"2", "add", "3"}),
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

func (s *calculationMultipleInputTestSuite) TestValidate() {
	input1 := "2"
	input2 := "3"
	testCases := []struct {
		Desc        string
		Input       calculation.ICalculationInput
		ExpectedErr error
	}{
		{
			Desc:        "success-validate-calculation-with-multiple-input",
			Input:       calculation.NewCalculationWithMultipleInput([]string{input1, "add", input2}),
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-validate-with-non-odd-input-length",
			Input:       calculation.NewCalculationWithMultipleInput([]string{"2", "add"}),
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-with-not-listed-operation",
			Input:       calculation.NewCalculationWithMultipleInput([]string{input1, "random", input2}),
			ExpectedErr: errors.ErrInvalidOperation,
		},
		{
			Desc:        "failed-validate-with-invalid-input-to-be-operated",
			Input:       calculation.NewCalculationWithMultipleInput([]string{input1, "add", "random"}),
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

func (s *calculationMultipleInputTestSuite) TestChangeToPostfixOperation() {
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
			Desc:        "success-change-multiple-inputs-with-equal-priority-operation-to-postfix-1",
			Inputs:      []string{"1", "add", "3", "subtract", "5"},
			ExpectedRes: []string{"1", "3", "add", "5", "subtract"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-equal-priority-operation-to-postfix-2",
			Inputs:      []string{"1", "multiply", "3", "divide", "5"},
			ExpectedRes: []string{"1", "3", "multiply", "5", "divide"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-to-postfix-1",
			Inputs:      []string{"1", "add", "3", "multiply", "5"},
			ExpectedRes: []string{"1", "3", "5", "multiply", "add"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-to-postfix-2",
			Inputs:      []string{"1", "multiply", "3", "add", "5"},
			ExpectedRes: []string{"1", "3", "multiply", "5", "add"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-to-postfix-3",
			Inputs:      []string{"1", "add", "3", "multiply", "5", "subtract", "2"},
			ExpectedRes: []string{"1", "3", "5", "multiply", "add", "2", "subtract"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-to-postfix-4",
			Inputs:      []string{"1", "multiply", "3", "add", "5", "divide", "2"},
			ExpectedRes: []string{"1", "3", "multiply", "5", "2", "divide", "add"},
		},
		{
			Desc:        "success-change-multiple-inputs-with-non-equal-priority-operation-to-postfix-5",
			Inputs:      []string{"1", "add", "3", "multiply", "5", "divide", "5"},
			ExpectedRes: []string{"1", "3", "5", "multiply", "5", "divide", "add"},
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculationWithMultipleInput := calculation.NewCalculationWithMultipleInput(tC.Inputs)
			res := calculation.GetChangeToPostfixOperation(&calculationWithMultipleInput, tC.Inputs)

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}

func (s *calculationMultipleInputTestSuite) TestCalculatePostfixOperation() {
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
		{
			Desc:        "success-calculate-multiple-inputs-with-non-equal-priority-postfix-operation-5",
			Inputs:      []string{"1", "3", "5", "multiply", "5", "divide", "add"},
			ExpectedRes: "4",
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculationWithMultipleInput := calculation.NewCalculationWithMultipleInput(tC.Inputs)
			res := calculation.GetCalculatePostfixOperation(&calculationWithMultipleInput, tC.Inputs)

			require.Equal(t, tC.ExpectedRes, res)
		})
	}
}
