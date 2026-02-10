package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/server"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env con error: ", err)
	}

	db.InitDB()
	r := server.SetupRoutes()

	r.Run("Localhost:8080")
	fmt.Printf("(\"Servidor corriendo en http://localhost:8080\")/n")
}
