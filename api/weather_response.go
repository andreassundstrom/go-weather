package api

type WeatherResponse struct {
	ApprovedTime  string `json:"approvedTime"`
	ReferenceTime string `json:"referenceTime"`
	TimeSeries    []struct {
		ValidTime  string `json:"validTime"`
		Parameters []struct {
			Name      string    `json:"name"`
			LevelType string    `json:"levelType"`
			Level     int       `json:"level"`
			Unit      string    `json:"unit"`
			Values    []float32 `json:"values"`
		} `json:"parameters"`
	} `json:"timeSeries"`
}
