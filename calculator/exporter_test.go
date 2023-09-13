package calculator

var (
	GetParseInput                                   = (*calculatorService).parseInput
	GetGetCalculationType                           = (*calculatorService).getCalculationType
	GetValidateAndConstructCalculationOneInput      = (*calculatorService).validateAndConstructCalculationOneInput
	GetDoCalculationWithOneInput                    = (*calculatorService).doCalculationWithOneInput
	GetValidateAndConstructCalculationMultipleInput = (*calculatorService).validateAndConstructCalculationMultipleInput
	GetDoCalculationWithMultipleInput               = (*calculatorService).doCalculationWithMultipleInput
)
