package cmd

import (
	"andreassundstrom/go-weather/api"
	"fmt"

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

	for _, timeSeries := range weatherResponse.TimeSeries {
		for _, parameter := range timeSeries.Parameters {
			if parameter.Name == "t" {
				fmt.Printf("%s: %f\n", timeSeries.ValidTime, parameter.Values[0])
			}
		}
	}
}
