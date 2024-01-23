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

type weatherData struct {
	maxTemp       float32
	precipitation float32
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

	forecastData := make(map[string]weatherData)

	for _, timeSeries := range weatherResponse.TimeSeries {

		parsedTime, err := time.Parse(time.RFC3339, timeSeries.ValidTime)
		if err != nil {
			fmt.Println("Failed to parse time: ", err)
		}
		date := parsedTime.Local().Format(time.DateOnly)

		currentWeatherData := weatherData{
			maxTemp:       timeSeries.GetParameter(api.Temperature),
			precipitation: timeSeries.GetParameter(api.PrecipitationMean),
		}

		if _, exists := forecastData[date]; exists {
			if forecastData[date].maxTemp < currentWeatherData.maxTemp {
				forecastData[date] = weatherData{
					maxTemp:       currentWeatherData.maxTemp,
					precipitation: forecastData[date].precipitation,
				}
			}

			if forecastData[date].precipitation < currentWeatherData.precipitation {
				forecastData[date] = weatherData{
					maxTemp:       forecastData[date].maxTemp,
					precipitation: currentWeatherData.precipitation,
				}
			}
		} else {
			forecastData[date] = currentWeatherData
		}
	}

	/* Sort the dates */
	dates := make([]string, 0, len(forecastData))
	for k := range forecastData {
		dates = append(dates, k)
	}
	sort.Strings(dates)

	/* Print the dates */
	fmt.Println("============================================")
	fmt.Println("Date\t\tMaxTemp\tPrecipitation (mm/h)")
	fmt.Println("============================================")
	for _, date := range dates {

		fmt.Println(fmt.Sprintf("%s\t%.0f°C\t%.0f", date, forecastData[date].maxTemp, forecastData[date].precipitation))
	}
}
