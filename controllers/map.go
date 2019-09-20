package controllers

import (
	"net/http"
	"planets/weather"
	"strconv"
)

func MapHandler(w http.ResponseWriter, r *http.Request) {
	dayString := r.URL.Query().Get("dia")
	day, _ := strconv.Atoi(dayString)

	scaleString := r.URL.Query().Get("escala")
	scale, _ := strconv.Atoi(scaleString)

	weatherMap := weather.DrawMap(day, scale)
	w.Write([]byte(weatherMap))
}
