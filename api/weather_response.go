package api

const Temperature = "t"
const PrecipitationMean = "pmean"

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

func (timeSeries *TimeSerie) GetParameter(parameterType string) float32 {

	for _, parameter := range timeSeries.Parameters {
		if parameter.Name == parameterType {
			return parameter.Values[0]
		}
	}
	return 0.
}
