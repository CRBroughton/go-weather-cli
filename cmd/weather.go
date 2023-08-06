package weather

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get the current weather",
	Long:  "Gets the current weather for the specified location",

	Run: func(cmd *cobra.Command, args []string) {
		println("hello world")
	},
}

func Execute() {
	if err := weatherCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
