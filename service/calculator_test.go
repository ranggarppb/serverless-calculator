package calculator_test

import (
	"testing"

	calculator "github.com/ranggarppb/serverless-calculator/service"
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
	testCases := []struct {
		Desc        string
		Input       string
		ExpectedRes []string
		ExpectedErr error
	}{
		{
			Desc:        "success-parsing-with-normal-input",
			Input:       "1 add 2",
			ExpectedRes: []string{"1", "add", "2"},
			ExpectedErr: nil,
		},
	}

	for _, tC := range testCases {
		s.T().Run(tC.Desc, func(t *testing.T) {
			calculatorSvc := calculator.NewCalculatorService()
			res, err := calculator.GetParseInput(calculatorSvc, tC.Input)

			if tC.ExpectedErr != nil {
				require.Equal(t, tC.ExpectedErr.Error(), err.Error())
			} else {
				require.Empty(t, err)
				require.Equal(t, tC.ExpectedRes, res)
			}
		})
	}
}
