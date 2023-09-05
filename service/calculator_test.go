package calculator_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type calculatorTestSuite struct {
	suite.Suite
}

func TestSuiteCalculator(t *testing.T) {
	s := calculatorTestSuite{}
	suite.Run(t, &s)
}
