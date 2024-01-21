package main

import (
	"andreassundstrom/go-weather/api"
	"flag"
	"fmt"
)

func main() {
	actionPtr := flag.String("command", "not set", "a string")
	flag.Parse()
	if *actionPtr == "weather" {
		weatherResponse, error := api.GetWeather(59.334591, 18.063240)

		if error != nil {
			println("Could not get weather data")
			return
		}

		println("Got result: ")
		for _, timeSeries := range weatherResponse.TimeSeries {
			for _, parameter := range timeSeries.Parameters {
				if parameter.Name == "t" {
					fmt.Printf("%s: %f\n", timeSeries.ValidTime, parameter.Values[0])
				}
			}
		}
	}
}
