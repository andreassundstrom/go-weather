package cmd

import (
	"andreassundstrom/go-weather/api"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(forecastCmd)
}

type weatherData struct {
	maxTemp       float32
	precipitation float32
	symbol        float32
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

		date := timeSeries.GetValidDate()

		currentWeatherData := weatherData{
			maxTemp:       timeSeries.GetParameter(api.Temperature),
			precipitation: timeSeries.GetParameter(api.PrecipitationMean),
			symbol:        timeSeries.GetParameter(api.WeatherSymbol),
		}

		if _, exists := forecastData[date]; exists {
			if forecastData[date].maxTemp < currentWeatherData.maxTemp {
				forecastData[date] = weatherData{
					maxTemp:       currentWeatherData.maxTemp,
					precipitation: forecastData[date].precipitation,
					symbol:        forecastData[date].symbol,
				}
			}

			if forecastData[date].precipitation < currentWeatherData.precipitation {
				forecastData[date] = weatherData{
					maxTemp:       forecastData[date].maxTemp,
					precipitation: currentWeatherData.precipitation,
					symbol:        forecastData[date].symbol,
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
	fmt.Println("========================================================")
	fmt.Println("Date\t\tMaxTemp\tPrecipitation (mm/h)\tSymbol")
	fmt.Println("========================================================")
	for _, date := range dates {

		fmt.Println(fmt.Sprintf("%s\t%.0f°C\t%.0f\t\t\t%s", date, forecastData[date].maxTemp, forecastData[date].precipitation, getForecastSymbol(forecastData[date].symbol)))
	}
}

func getForecastSymbol(symbol float32) string {
	switch symbol {
	case 1.:
		return "☀️"
	case 2.:
		return "☀️"
	case 3.:
		return "🌤️"
	case 4.:
		return "⛅"
	case 5.:
		return "☁️"
	case 6.:
		return "☁️"
	// Light rain
	case 18.:
		return "☔"
	// Moderate rain
	case 19.:
		return "☔☔"
	// Heavy rain
	case 20.:
		return "☔☔☔"
	// Light snowfall
	case 25.:
		return "❄️"
	// Moderate snowfall
	case 26.:
		return "❄️❄️"
	// Heavy snowfall
	case 27.:
		return "❄️❄️❄️"
	default:
		return fmt.Sprintf("%.0f", symbol)
	}
}
