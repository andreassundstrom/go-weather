package cmd

import (
	"andreassundstrom/go-weather/api"
	"fmt"
	"sort"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(forecastCmd)
}

var forecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get current weather forecast",
	Run: func(cmd *cobra.Command, args []string) {
		forecast()
	},
}

func forecast() {
	weatherResponse, error := api.GetWeather(59.334591, 18.063240)

	if error != nil {
		println("Failed to get weather data")
		return
	}

	forecastData := make(map[string]float32)

	for _, timeSeries := range weatherResponse.TimeSeries {
		for _, parameter := range timeSeries.Parameters {
			if parameter.Name == "t" {
				parsedTime, err := time.Parse(time.RFC3339, timeSeries.ValidTime)
				if err != nil {
					fmt.Println("Failed to parse time: ", err)
				}
				date := parsedTime.Local().Format(time.DateOnly)
				newValue := parameter.Values[0]
				if currentValue, ok := forecastData[date]; ok {
					if currentValue < newValue {
						forecastData[date] = newValue
					}
				} else {
					forecastData[date] = newValue
				}
			}
		}
	}

	/* Sort the dates */
	dates := make([]string, 0, len(forecastData))
	for k := range forecastData {
		dates = append(dates, k)
	}
	sort.Strings(dates)

	/* Print the dates */
	for _, date := range dates {
		fmt.Println(fmt.Sprintf("%s: %.0f°C", date, forecastData[date]))
	}
}
