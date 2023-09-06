package calculator

import "strings"

type calculatorService struct{}

func NewCalculatorService() *calculatorService {
	return &calculatorService{}
}

func (c *calculatorService) GetCalculationHistory() []string {
	result := []string{
		"Hello World!",
		"hello World!",
	}

	return result
}

func (c *calculatorService) Calculate(input string) (string, error) {

	return input, nil
}

func (c *calculatorService) parseInput(input string) ([]string, error) {
	return strings.Split(input, " "), nil
}
