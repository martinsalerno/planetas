package models

type Weather struct {
	Day       int     `gorm:"primary_key" json:"dia"`
	Climate   string  `gorm:"not null" json:"clima"`
	Intensity float64 `gorm:"not null; default: 0.0" json:"-"`
}
