package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ranggarppb/serverless-calculator/calculator"
	"github.com/ranggarppb/serverless-calculator/errors"
	"github.com/ranggarppb/serverless-calculator/utils"
	"github.com/spf13/cobra"
)

var consoleApp = &cobra.Command{
	Use:   "console",
	Short: "Start console application",
	Run:   startConsole,
}

func init() {
	rootCmd.AddCommand(consoleApp)
}

func startConsole(cmd *cobra.Command, args []string) {
	var input string
	var res string = "0"
	var calculationResult calculator.CalculatorResult
	var calculationHistory calculator.CalculationHistory
	var err errors.WrappedError
	consoleReader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	fmt.Print("Enter operation\n")

Operation:
	for {
		input, _ = consoleReader.ReadString('\n')
		trimmedInput := strings.TrimSuffix(input, "\n")

		switch trimmedInput {
		case utils.COMMAND_EXIT:
			break Operation
		case utils.COMMAND_CANCEL:
			res = "0"
			fmt.Println(res)
		case utils.COMMAND_HISTORY:
			fmt.Println(calculationHistory.History)
		default:
			calculationResult, err = doOperation(ctx, trimmedInput, res, calculationHistory)

			if err != nil {
				fmt.Printf("Get error %s, %s\n", err.ErrCode(), err.Error())
			} else {
				fmt.Println(calculationResult.Result)
				calculationHistory.History = append(calculationHistory.History, calculationResult)
				res = calculationResult.Result
			}
		}
	}
}

func doOperation(ctx context.Context, trimmedInput string, initial string, history calculator.CalculationHistory) (calculator.CalculatorResult, errors.WrappedError) {
	inputs := strings.Split(trimmedInput, " ")
	switch {
	case utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, inputs[0]):
		return calculatorService.Calculate(ctx, fmt.Sprintf("%s %s", trimmedInput, initial))
	case inputs[0] == utils.REPEAT:
		repeatInput, err := validateRepeatOperation(inputs, history)

		if err != nil {
			return calculator.CalculatorResult{}, err
		}

		return doRepeatOperation(ctx, initial, repeatInput, history)

	default:
		return calculatorService.Calculate(ctx, fmt.Sprintf("%s %s", initial, trimmedInput))
	}
}

func validateRepeatOperation(inputs []string, history calculator.CalculationHistory) (int, errors.WrappedError) {
	if len(inputs) > 2 {
		return 0, errors.ErrInvalidOperation
	}
	repeatInput, err := strconv.Atoi(inputs[1])

	if err != nil || repeatInput > len(history.History) {
		return 0, errors.ErrInvalidInputToBeOperated
	}

	return repeatInput, nil
}

func doRepeatOperation(ctx context.Context, initial string, repeatInput int, history calculator.CalculationHistory) (
	calculator.CalculatorResult, errors.WrappedError) {
	operationToBeRepeated := strings.Split(history.History[len(history.History)-repeatInput].Input, " ")

	if utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, operationToBeRepeated[0]) {
		return calculatorService.Calculate(ctx, fmt.Sprintf("%s %s", operationToBeRepeated[0], initial))
	} else {
		operationToBeRepeated[0] = initial
		return calculatorService.Calculate(ctx, strings.Join(operationToBeRepeated, " "))
	}
}
