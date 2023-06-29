package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mounakh_cli/common"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	accessKey := os.Getenv("WEATHER_API_KEY")

	if len(os.Args) >= 2 {
		common.Q = os.Args[1] + ", France"
	}

	fullURL := &url.URL{
		Scheme:   "http",
		Host:     "api.weatherstack.com",
		Path:     common.Path.Route,
		RawQuery: "access_key=" + accessKey + "&query=" + common.Q,
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

	var weather common.Weather
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

	if current.WeatherDescriptions[0] == string(common.Sunny) {
		color.Yellow(message)
	} else {
		fmt.Println(message)
	}
}

func Hello() string {
	return "hello"
}
