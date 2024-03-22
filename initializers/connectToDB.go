package initializers

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ti *time.Location

func ConnectTODBPostgres() {
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

func ConnectTODBMSSQL() {
	var err error

	dsn := os.Getenv("DSN")
	DB, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Fallo en conexion a base de datos...")
	}
}
