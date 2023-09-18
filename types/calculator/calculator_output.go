package calculator

type CalculationResult struct {
	Input  string `json:"input"`
	Result string `json:"result"`
}

type CalculationHistory struct {
	History []CalculationResult `json:"result"`
}
