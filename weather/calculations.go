package weather

import (
	"planets/models"
)

func PopulateWeatherTable() {
	planets := []models.Planet{}

	models.DB.Find(&planets)

	for year := 0; year < 10; year++ {
		startDay := year * 365
		endDay := startDay + 365

		go populateRangeWeather(startDay, endDay, planets)
	}
}

func CalculateDayWeather(day int) models.Weather {
	planets := []models.Planet{}

	models.DB.Find(&planets)

	triangle := models.Triangle{
		Vertex1: planets[0].Position(day),
		Vertex2: planets[1].Position(day),
		Vertex3: planets[2].Position(day),
	}

	area, pointsAligned, includesSun := triangle.Stats()

	weather := models.Weather{
		Day:       day,
		Climate:   climateForConditions(pointsAligned, includesSun),
		Intensity: area,
	}

	return weather
}

func populateRangeWeather(startDay int, endDay int, planets []models.Planet) {
	for day := startDay; day < endDay; day++ {
		weather := CalculateDayWeather(day)

		models.DB.Create(&weather)
	}
}

func climateForConditions(pointsAligned bool, includesSun bool) string {
	if pointsAligned && includesSun {
		return "sequia"
	} else if pointsAligned && !includesSun {
		return "condiciones optimas de presión y temperatura"
	} else if !pointsAligned && includesSun {
		return "lluvia"
	} else {
		return "indeterminado"
	}
}
