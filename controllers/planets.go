package controllers

import (
	"encoding/json"
	"net/http"
	"planets/models"
	"strconv"
)

func PlanetsHandler(w http.ResponseWriter, r *http.Request) {
	dayString := r.URL.Query().Get("dia")
	day, _ := strconv.Atoi(dayString)

	planets := []models.Planet{}
	points := []models.Point{}
	models.DB.Find(&planets)

	for _, planet := range planets {
		points = append(points, planet.Position(day))
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(&points)
}
