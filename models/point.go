package models

type Point struct {
	PlanetName string  `json:"planeta"`
	Angle      float64 `json:"angulo"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
}
