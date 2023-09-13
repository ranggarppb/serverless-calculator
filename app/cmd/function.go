package cmd

import (
	"context"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	f "github.com/ranggarppb/serverless-calculator/internal/function"
	"github.com/spf13/cobra"
)

var functionCommand = &cobra.Command{
	Use:   "function",
	Short: "Start HTTP function",
	Run:   function,
}

func init() {
	rootCmd.AddCommand(functionCommand)
}

func function(cmd *cobra.Command, args []string) {
	// Use PORT environment variable, or default to 8080.
	port := "8080"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	funcframework.RegisterHTTPFunctionContext(context.TODO(), "./", f.CreateCalculateFunction(calculatorRestHandler))

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
