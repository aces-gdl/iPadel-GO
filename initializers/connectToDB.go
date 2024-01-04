package initializers

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ti *time.Location

func ConnectTODB() {
	var err error
	dsn := os.Getenv("DSN")
	ti, _ = time.LoadLocation("America/Mexico_City")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().In(ti)
		},
	})

	if err != nil {
		panic("Fallo en conexion a base de datos...")
	}
}
