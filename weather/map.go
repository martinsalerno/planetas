package weather

import (
	"math"
	"planets/constants"
	"planets/models"
	"strings"
)

func DrawMap(day int, scale int) string {
	if scale < constants.DefaultScaleThreshold {
		scale = constants.DefaultScale
	}

	planets := []models.Planet{}
	models.DB.Order("distance_from_sun desc").Find(&planets)

	matrix := drawMatrix(scale)

	placePlanets(planets, day, scale, matrix)

	aux := []string{}

	for _, vector := range matrix {
		aux = append(aux, strings.Join(vector, " "))
	}

	return strings.Join(aux, "\n")
}

func placePlanets(planets []models.Planet, day int, scale int, matrix [][]string) {
	furthestPoint := planets[0].DistanceFromSun

	position1 := scalePoint(planets[0].Position(day), scale, furthestPoint)
	position2 := scalePoint(planets[1].Position(day), scale, furthestPoint)
	position3 := scalePoint(planets[2].Position(day), scale, furthestPoint)

	matrix[int(position1.Y)][int(position1.X)] = string(position1.PlanetName[0])
	matrix[int(position2.Y)][int(position2.X)] = string(position2.PlanetName[0])
	matrix[int(position3.Y)][int(position3.X)] = string(position3.PlanetName[0])
}

func drawMatrix(scale int) [][]string {
	matrix := initializeMatrix(scale)

	for i, vector := range matrix {
		for j := range vector {
			matrix[i][j] = "."
			matrix[scale-1][j] = "-"
		}
		matrix[i][scale-1] = "|"
	}
	matrix[scale-1][scale-1] = "O"

	return matrix
}

func initializeMatrix(scale int) [][]string {
	matrix := make([][]string, scale *2 + 1)
	for i := range matrix {
		matrix[i] = make([]string, scale *2 + 1)
	}

	return matrix
}

func scalePoint(point models.Point, scale int, furthestPoint float64) models.Point {
	scaleFactor := furthestPoint / float64(scale)

	return models.Point{
		X:          math.Ceil(((point.X + furthestPoint) / scaleFactor)),
		Y:          math.Ceil(((point.Y - furthestPoint) / -scaleFactor)),
		PlanetName: point.PlanetName,
	}
}
