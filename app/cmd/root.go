package cmd

import (
	"log"

	"github.com/ranggarppb/serverless-calculator/calculator"
	"github.com/ranggarppb/serverless-calculator/internal/rest"
	cl "github.com/ranggarppb/serverless-calculator/objects/calculation"
	"github.com/spf13/cobra"
)

var (
	calculatorService     cl.ICalculatorService
	calculatorRestHandler cl.ICalculatorRestHandler
)

var (
	rootCmd = &cobra.Command{
		Use:   "serverlesscalculator",
		Short: "serverlesscalculator is application for running calculator in serverless and in local",
	}
)

// Execute will call the root command execute
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(initApp)
}

func initApp() {
	calculatorService = calculator.NewCalculatorService()
	calculatorRestHandler = rest.NewCalculatorRestHandler(calculatorService)
}
