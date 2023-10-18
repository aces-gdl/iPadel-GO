package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	fmt.Println("...Inicializando")

}
