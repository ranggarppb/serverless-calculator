package utils

const (
	ADDITION       = "add"
	SUBSTRACTION   = "substract"
	MULTIPLICATION = "multiply"
	DIVISION       = "divide"
	NEGATION       = "neg"
	SQUARE         = "sqr"
	SQUAREROOT     = "sqrt"
	ABSOLUTE       = "abs"
)

var OPERATIONS_WITH_ONE_INPUTS = []string{NEGATION, SQUARE, SQUAREROOT, ABSOLUTE}

var OPERATIONS_WITH_TWO_INPUTS = []string{ADDITION, SUBSTRACTION, MULTIPLICATION, DIVISION}
