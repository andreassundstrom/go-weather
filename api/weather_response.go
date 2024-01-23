package api

import (
	"fmt"
	"time"
)

const Temperature = "t"
const PrecipitationMean = "pmean"
const WeatherSymbol = "Wsymb2"

type WeatherResponse struct {
	ApprovedTime  string      `json:"approvedTime"`
	ReferenceTime string      `json:"referenceTime"`
	TimeSeries    []TimeSerie `json:"timeSeries"`
}

type TimeSerie struct {
	ValidTime  string      `json:"validTime"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Name      string    `json:"name"`
	LevelType string    `json:"levelType"`
	Level     int       `json:"level"`
	Unit      string    `json:"unit"`
	Values    []float32 `json:"values"`
}

func (timeSerie *TimeSerie) GetValidDate() string {
	parsedTime, err := time.Parse(time.RFC3339, timeSerie.ValidTime)
	if err != nil {
		fmt.Println("Failed to parse time: ", err)
	}
	return parsedTime.Local().Format(time.DateOnly)
}

func (timeSeries *TimeSerie) GetParameter(parameterType string) float32 {

	for _, parameter := range timeSeries.Parameters {
		if parameter.Name == parameterType {
			return parameter.Values[0]
		}
	}
	panic(fmt.Sprintf("Failed to get parameter of type %s", parameterType))
}
