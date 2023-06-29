package main

import (
	"encoding/json"
	"fmt"

	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

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

const (
	Sunny  WeatherCondition = "Sunny"
	Cloudy WeatherCondition = "Cloudy"
	Rainy  WeatherCondition = "Rainy"
	Snowy  WeatherCondition = "Snowy"
	Windy  WeatherCondition = "Windy"
)

var q = "Lille"

var path = URL{
	Route: "/current",
	Query: "Lille, France",
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	accessKey := os.Getenv("WEATHER_API_KEY")

	if len(os.Args) >= 2 {
		q = os.Args[1] + ", France"
	}

	fullURL := &url.URL{
		Scheme:   "http",
		Host:     "api.weatherstack.com",
		Path:     path.Route,
		RawQuery: "access_key=" + accessKey + "&query=" + q,
	}

	res, err := http.Get(fullURL.String())

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("WeatherAPI not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current := weather.Location, weather.Current

	currentTime := time.Now()
	currentTimeFormatted := currentTime.Format("15:04:05")

	message := fmt.Sprintf(
		"%s, %s : %.0f°C, %s, %s\n",
		location.Name,
		location.Country,
		current.Temperature,
		current.WeatherDescriptions[0],
		currentTimeFormatted,
	)

	if current.WeatherDescriptions[0] == string(Sunny) {
		color.Yellow(message)
	} else {
		fmt.Println(message)
	}
}

func Hello() string {
	return "hello"
}
