package models

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	dbFile, exists := os.LookupEnv("DATABASE_FILE")

	if !exists {
		log.Print("Missing DATABASE_FILE env variable")
	}

	var err error
	DB, err = gorm.Open("sqlite3", dbFile)

	if err != nil {
		log.Panic(err)
	}

	DB.AutoMigrate(&Planet{})
	DB.AutoMigrate(&Weather{})

  ferengi := Planet {
    Name: "Ferengi",
    DistanceFromSun: 500.0,
    AngularVelocity: 1.0,
    InitialDegrees: 90.0,
  }
  betasoide := Planet {
    Name: "Betasoide",
    DistanceFromSun: 2000.0,
    AngularVelocity: 2.0,
    InitialDegrees: 90.0,
  }
  vulcano := Planet {
    Name: "Vulcano",
    DistanceFromSun: 1000.0,
    AngularVelocity: - 5.0,
    InitialDegrees: 90.0,
  }

  DB.Create(&ferengi)
  DB.Create(&betasoide)
  DB.Create(&vulcano)
}

func CloseDB() error {
	return DB.Close()
}
