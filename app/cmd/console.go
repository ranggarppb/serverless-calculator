package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

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
	var result string
	var err errors.WrappedError
	consoleReader := bufio.NewReader(os.Stdin)
	ctx := context.Background()

	fmt.Print("Enter operation\n")

	for {
		input, _ = consoleReader.ReadString('\n')
		trimmedInput := strings.TrimSuffix(input, "\n")

		if trimmedInput == "exit" {
			break
		} else if trimmedInput == "cancel" {
			res = "0"
			fmt.Println(res)
		} else {
			if inputs := strings.Split(trimmedInput, " "); utils.ContainString(utils.OPERATIONS_WITH_ONE_INPUTS, inputs[0]) {
				result, err = calculatorService.Calculate(ctx, fmt.Sprintf("%s %s", trimmedInput, res))
			} else {
				result, err = calculatorService.Calculate(ctx, fmt.Sprintf("%s %s", res, trimmedInput))
			}
			if err != nil {
				fmt.Printf("Get error %s, %s\n", err.ErrCode(), err.Error())
			} else {
				fmt.Println(result)
				res = result
			}
		}
	}
}
