package types

type CalculatorInput struct {
	Input string `json:"input"`
}

type CalculatorResult struct {
	Input  string `json:"input"`
	Result string `json:"result"`
}

type CalculatorHistory struct {
	Result []string `json:"result"`
}

type ICalculatorService interface {
	GetCalculationHistory() []string
	Calculate(input string) string
}
