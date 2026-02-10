package main

import (
	"os"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/tu-usuario/blog-api/internal/auth"
	"github.com/tu-usuario/blog-api/internal/constants"
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

	email := os.Getenv(constants.EnvAdminEmail)
	password := os.Getenv(constants.EnvAdminPassword)
	err = auth.SeedAdmin(email, password)
	
	if err!=nil{
		log.Printf("❌ No se pudo crear el admin inicial: %v", err)
	}else{
		log.Printf("✅ Se cargo exitosamente el admin inicial")
	}

	r.Run("Localhost:8080")
	fmt.Printf("(\"Servidor corriendo en http://localhost:8080\")/n")
}
