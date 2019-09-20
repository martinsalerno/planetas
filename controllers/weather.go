package controllers

import (
	"planets/models"
	"planets/weather"
	"encoding/json"
	"net/http"
	"strconv"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	dayString := r.URL.Query().Get("dia")
	day, _ := strconv.Atoi(dayString)

	weatherModel := models.Weather{}
	models.DB.Where("day = ?", day).First(&weatherModel)
  
  if weatherModel.Day == 0 {
  	weatherModel = weather.CalculateDayWeather(day)
  }

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&weatherModel)
}
