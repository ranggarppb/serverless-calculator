package calculator

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

func (c *calculatorService) Calculate(input string) string {

	return input
}
