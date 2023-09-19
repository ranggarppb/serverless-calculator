package calculation_test

import (
	"testing"

	"github.com/ranggarppb/serverless-calculator/errors"
	cl "github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type calculationInputTestSuite struct {
	suite.Suite
}

func TestSuiteCalculationInput(t *testing.T) {
	s := calculationInputTestSuite{}
	suite.Run(t, &s)
}

func (s *calculationInputTestSuite) TestParseCalculationInput() {
	testCases := []struct {
		Desc        string
		Input       string
		History     cl.CalculationHistory
		ExpectedRes cl.ICalculationInput
		ExpectedErr error
	}{
		{
			Desc:        "success-parse-input-for-calculation-with-one-input-1",
			Input:       " sqr 2",
			ExpectedRes: cl.NewCalculationWithOneInput("sqr", []string{"2"}),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parse-input-for-calculation-with-one-input-2",
			Input:       " sqr 2 3",
			ExpectedRes: cl.NewCalculationWithOneInput("sqr", []string{"2", "3"}),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parse-input-for-calculation-with-one-input-2",
			Input:       "2 sqr 3",
			ExpectedRes: cl.NewCalculationWithOneInput("sqr", []string{"2", "3"}),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parse-input-for-calculation-with-multiple-input",
			Input:       "1 add 2",
			ExpectedRes: cl.NewCalculationWithMultipleInput([]string{"1", "add", "2"}),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parse-input-for-calculation-with-multiple-input-2",
			Input:       "1 add 2 multiply 3",
			ExpectedRes: cl.NewCalculationWithMultipleInput([]string{"1", "add", "2", "multiply", "3"}),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parse-input-for-calculation-with-multiple-input-3",
			Input:       "1 add random",
			ExpectedRes: cl.NewCalculationWithMultipleInput([]string{"1", "add", "random"}),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parse-input-for-calculation-with-repeat-operation-1",
			Input:       "2 repeat 2",
			History:     cl.CalculationHistory{History: []cl.CalculationResult{{Input: "1 add 3", Result: "4"}}},
			ExpectedRes: cl.NewCalculationRepeatInput([]string{"repeat", "2"}, "2", cl.CalculationHistory{History: []cl.CalculationResult{{Input: "1 add 3", Result: "4"}}}),
			ExpectedErr: nil,
		},
		{
			Desc:        "success-parse-input-for-calculation-with-repeat-operation-2",
			Input:       "2 repeat 2 3",
			History:     cl.CalculationHistory{History: []cl.CalculationResult{{Input: "1 add 3", Result: "4"}}},
			ExpectedRes: cl.NewCalculationRepeatInput([]string{"repeat", "2", "3"}, "2", cl.CalculationHistory{History: []cl.CalculationResult{{Input: "1 add 3", Result: "4"}}}),
			ExpectedErr: nil,
		},
		{
			Desc:        "failed-parse-input-with-empty-input",
			Input:       "",
			ExpectedRes: nil,
			ExpectedErr: errors.ErrInvalidInput,
		},
		{
			Desc:        "failed-parse-input-with-input-less-than-2-words",
			Input:       "sqr",
			ExpectedRes: nil,
			ExpectedErr: errors.ErrInvalidOperation,
		},
	}
	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculationInput := cl.CalculationInput{Input: tC.Input}
			res, err := calculationInput.ParseCalculationInput(tC.History)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr, err)
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}
