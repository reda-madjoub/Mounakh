package common

const (
	Sunny  WeatherCondition = "Sunny"
	Cloudy WeatherCondition = "Cloudy"
	Rainy  WeatherCondition = "Rainy"
	Snowy  WeatherCondition = "Snowy"
	Windy  WeatherCondition = "Windy"
)

var Q = "Lille"

var Path = URL{
	Route: "/current",
	Query: "Lille, France",
}
