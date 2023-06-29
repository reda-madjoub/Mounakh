package common

type URL struct {
	Route string
	Query string
}

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Country   string `json:"country"`
		TimeEpoch int64  `json:"localtime_epoch"`
	} `json:"location"`
	Current struct {
		Temperature         float64  `json:"temperature"`
		WeatherDescriptions []string `json:"weather_descriptions"`
		WeatherIcons        []string `json:"weather_icons"`
	} `json:"current"`
}

type WeatherCondition string
