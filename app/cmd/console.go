package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ranggarppb/serverless-calculator/errors"
	cl "github.com/ranggarppb/serverless-calculator/objects/calculation"
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
	var calculationResult cl.CalculationResult
	var calculationHistory cl.CalculationHistory
	var err errors.WrappedError
	consoleReader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	fmt.Print("List Commands\n")
	fmt.Println("a. Exiting session: exit")
	fmt.Println("b. Clear all result: cancel")
	fmt.Println("c. Showing calculation history: history")
	fmt.Println("d. Operation: ${single input operation} or ${multiple input operation} ${input} (see README Features section for more detail)")
	fmt.Println("")
	fmt.Print("Enter command (start from 0)\n")

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

func doOperation(ctx context.Context, trimmedInput string, initial string, history cl.CalculationHistory) (cl.CalculationResult, errors.WrappedError) {
	calculationInput := cl.CalculationInput{Input: fmt.Sprintf("%s %s", initial, trimmedInput)}

	parsedCalculationInput, err := calculationInput.ParseCalculationInput(history)

	if err != nil {
		return cl.CalculationResult{}, err
	}

	return calculatorService.Calculate(ctx, parsedCalculationInput)
}
