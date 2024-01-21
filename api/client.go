package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl string = "https://opendata-download-metfcst.smhi.se"

func GetWeather(lat float32, lon float32) (WeatherResponse, error) {

	var weatherResponse WeatherResponse

	url := fmt.Sprintf("%s/api/category/pmp3g/version/2/geotype/point/lon/%f/lat/%f/data.json", baseUrl, lon, lat)
	fmt.Printf("Getting weather data from %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		return weatherResponse, nil
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if json.Unmarshal(body, &weatherResponse) != nil {
		fmt.Println("Failed to parse weather response as json")
	}
	return weatherResponse, nil
}
