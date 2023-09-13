package cmd

import (
	"log"

	calculator "github.com/ranggarppb/serverless-calculator/service"
	"github.com/ranggarppb/serverless-calculator/types"
	"github.com/spf13/cobra"
)

var (
	calculatorService types.ICalculatorService
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
}
