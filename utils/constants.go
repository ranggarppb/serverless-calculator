package utils

const (
	ADDITION        = "add"
	SUBSTRACTION    = "subtract"
	MULTIPLICATION  = "multiply"
	DIVISION        = "divide"
	NEGATION        = "neg"
	SQUARE          = "sqr"
	SQUAREROOT      = "sqrt"
	ABSOLUTE        = "abs"
	CUBE            = "cube"
	CUBERT          = "cubert"
	REPEAT          = "repeat"
	COMMAND_EXIT    = "exit"
	COMMAND_CANCEL  = "cancel"
	COMMAND_HISTORY = "history"
)

var OPERATIONS_WITH_ONE_INPUTS = []string{NEGATION, SQUARE, SQUAREROOT, ABSOLUTE, CUBE, CUBERT}

var OPERATIONS_WITH_MULTIPLE_INPUTS = []string{ADDITION, SUBSTRACTION, MULTIPLICATION, DIVISION}

var OPERATIONS_PRIORITY = map[string]int{ADDITION: 0, SUBSTRACTION: 0, MULTIPLICATION: 1, DIVISION: 1}
