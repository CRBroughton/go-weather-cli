package weather

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get the current weather",
	Long:  "Gets the current weather for the specified location",

	Run: func(cmd *cobra.Command, args []string) {
		getCurrentWeather()
	},
}

func getCurrentWeather() (weather []byte, err error) {
	url := "https://api.openweathermap.org/data/3.0/onecall?lat=50.8224243265497&lon=-0.13670887945602617&exclude=hourly,daily&appid=X"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return []byte(""), err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error reading weather response", err)
		return []byte(""), err
	}
	return body, nil
}

func Execute() {
	if err := weatherCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
