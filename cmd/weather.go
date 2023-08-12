package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/crbroughton/go-weather-cli/env"
	"github.com/spf13/cobra"
)

type WeatherDescription struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type CurrentWeather struct {
	Dt         int                  `json:"dt"`
	Sunrise    int                  `json:"sunrise"`
	Sunset     int                  `json:"sunset"`
	Temp       float64              `json:"temp"`
	FeelsLike  float64              `json:"feels_like"`
	Pressure   int                  `json:"pressure"`
	Humidity   int                  `json:"humidity"`
	DewPoint   float64              `json:"dew_point"`
	Uvi        float64              `json:"uvi"`
	Clouds     int                  `json:"clouds"`
	Visibility int                  `json:"visibility"`
	WindSpeed  float64              `json:"wind_speed"`
	WindDeg    int                  `json:"wind_deg"`
	WindGust   float64              `json:"wind_gust"`
	Weather    []WeatherDescription `json:"weather"`
}

type MinutelyData struct {
	Dt            int     `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

type HourlyData struct {
	Dt         int                  `json:"dt"`
	Temp       float64              `json:"temp"`
	FeelsLike  float64              `json:"feels_like"`
	Pressure   int                  `json:"pressure"`
	Humidity   int                  `json:"humidity"`
	DewPoint   float64              `json:"dew_point"`
	Uvi        float64              `json:"uvi"`
	Clouds     int                  `json:"clouds"`
	Visibility int                  `json:"visibility"`
	WindSpeed  float64              `json:"wind_speed"`
	WindDeg    int                  `json:"wind_deg"`
	WindGust   float64              `json:"wind_gust"`
	Weather    []WeatherDescription `json:"weather"`
	Pop        float64              `json:"pop"`
}

type TemperatureInfo struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type FeelsLikeInfo struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

type DailyData struct {
	Dt        int                  `json:"dt"`
	Sunrise   int                  `json:"sunrise"`
	Sunset    int                  `json:"sunset"`
	Moonrise  int                  `json:"moonrise"`
	Moonset   int                  `json:"moonset"`
	MoonPhase float64              `json:"moon_phase"`
	Summary   string               `json:"summary"`
	Temp      TemperatureInfo      `json:"temp"`
	FeelsLike FeelsLikeInfo        `json:"feels_like"`
	Pressure  int                  `json:"pressure"`
	Humidity  int                  `json:"humidity"`
	DewPoint  float64              `json:"dew_point"`
	WindSpeed float64              `json:"wind_speed"`
	WindDeg   int                  `json:"wind_deg"`
	WindGust  float64              `json:"wind_gust"`
	Weather   []WeatherDescription `json:"weather"`
	Clouds    int                  `json:"clouds"`
	Pop       float64              `json:"pop"`
	Rain      float64              `json:"rain"`
	Uvi       float64              `json:"uvi"`
}

type Alert struct {
	SenderName  string   `json:"sender_name"`
	Event       string   `json:"event"`
	Start       int      `json:"start"`
	End         int      `json:"end"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type WeatherResponse struct {
	Lat            float64        `json:"lat"`
	Lon            float64        `json:"lon"`
	Timezone       string         `json:"timezone"`
	TimezoneOffset int            `json:"timezone_offset"`
	Current        CurrentWeather `json:"current"`
	Minutely       []MinutelyData `json:"minutely"`
	Hourly         []HourlyData   `json:"hourly"`
	Daily          []DailyData    `json:"daily"`
	Alerts         []Alert        `json:"alerts"`
}

var tempType = "metric"
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Get the current weather",
	Long:  "Gets the current weather for the specified location",

	Run: func(cmd *cobra.Command, args []string) {
		weather, err := getCurrentWeather("https://api.openweathermap.org/data/3.0/onecall?lat=" + env.GetLat() + "lon=" + env.GetLon() + "&exclude=hourly,daily&units=" + tempType + "&appid=" + env.GetAPIKey())

		if err != nil {
			fmt.Println("Error getting current weather", err)
			return
		}

		fmt.Println("Current weather:", weather.Current.Weather[0].Main+" - "+weather.Current.Weather[0].Description)
		fmt.Println("Cloud cover:", weather.Current.Clouds)
		fmt.Println("Temperature:", weather.Current.Temp)
	},
}

func getCurrentWeather(url string) (weather WeatherResponse, err error) {
	var weatherResponse WeatherResponse

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return weatherResponse, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		fmt.Println("Error decoding weather response", err)
		return weatherResponse, err
	}

	return weatherResponse, nil
}

func Execute() {
	weatherCmd.Flags().StringVarP(&tempType, "type", "t", "metric", "Celcius or Fahrenheit")
	if err := weatherCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
