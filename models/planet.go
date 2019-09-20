package models

import (
	"github.com/jinzhu/gorm"
	"math"
	"planets/constants"
)

type Planet struct {
	gorm.Model
	Name            string  `gorm:"not null; unique"`
	DistanceFromSun float64 `gorm:"not null"`
	AngularVelocity float64 `gorm:"not null"`
	InitialDegrees  float64 `gorm:"not null; default: 90.0"`
}

func (planet Planet) angleFromSun(day int) float64 {
	totalDegreesTraveled := float64(day) * planet.AngularVelocity
	currentDegrees := math.Mod(planet.InitialDegrees-totalDegreesTraveled, constants.CircleDegrees)

	if math.Signbit(currentDegrees) {
		return constants.CircleDegrees + currentDegrees
	}

	return currentDegrees
}

func (planet Planet) Position(day int) Point {
	degreesFromSun := planet.angleFromSun(day)
	radiansFromSun := degreesFromSun * math.Pi / 180

	return Point{
		PlanetName: planet.Name,
		Angle:      math.Floor(degreesFromSun),
		X:          math.Round(planet.DistanceFromSun * math.Cos(radiansFromSun)),
		Y:          math.Round(planet.DistanceFromSun * math.Sin(radiansFromSun)),
	}
}
