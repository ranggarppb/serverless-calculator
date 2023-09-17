package calculator

var (
	GetParseInput                                   = (*calculatorService).parseInput
	GetGetCalculationType                           = (*calculatorService).getCalculationType
	GetValidateAndConstructCalculationOneInput      = (*calculatorService).validateAndConstructCalculationOneInput
	GetDoCalculationWithOneInput                    = (*calculatorService).doCalculationWithOneInput
	GetValidateAndConstructCalculationMultipleInput = (*calculatorService).validateAndConstructCalculationMultipleInput
	GetDoCalculationWithTwoInput                    = (*calculatorService).doCalculationWithTwoInput
	GetChangeToPostfixOperation                     = (*calculatorService).changeToPostfixOperation
	GetCalculatePostfixOperation                    = (*calculatorService).calculatePostfixOperation
)
