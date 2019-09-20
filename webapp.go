package main

import (
  "planets/controllers"
  "planets/models"
  "planets/weather"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strconv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return
	}

	populateWeather, exists := os.LookupEnv("POPULATE_WEATHER")
	populate, _ := strconv.ParseBool(populateWeather)

	if exists && populate {
		go weather.PopulateWeatherTable()
	}
}

func main() {
	defer models.CloseDB()

	r := mux.NewRouter()
	r.HandleFunc("/clima", controllers.WeatherHandler).Methods("GET").Queries("dia", "{[0-9]*?}")
	r.HandleFunc("/planetas", controllers.PlanetsHandler).Methods("GET").Queries("dia", "{[0-9]*?}")
	r.HandleFunc("/mapa", controllers.MapHandler).Methods("GET").Queries("dia", "{[0-9]*?}")

	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("Missing PORT env variable")
	}

	fmt.Printf("Server listening at %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
