package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tu-usuario/blog-api/internal/auth"
	"github.com/tu-usuario/blog-api/internal/constants"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/server"
)

func main() {

	loadEnv()

	initDB()
	seedAdmin()

	r := server.SetupRoutes()

	r.Run("Localhost:8080")
	fmt.Printf("(\"Servidor corriendo en http://localhost:8080\")/n")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error cargando el archivo .env con error: ", err)
	}
}


func initDB(){

	err := db.InitDB()

	if err!= nil{
		switch{
			case errors.Is(err, db.ErrFailedToConnectToDataBase):
				log.Fatal("❌ No se pudo conectar a MongoDB")
			case errors.Is(err, db.ErrFailledToCreateIndexis):
				log.Fatal("❌ Failed making emails indexis")
			default:
				log.Fatalf("❌ Unknown error: %v /n", err)
		}
	}

	log.Println("✅ Conectado a MongoDB exitosamente")
}

func seedAdmin() {

	email := os.Getenv(constants.EnvAdminEmail)
	password := os.Getenv(constants.EnvAdminPassword)
	err := auth.SeedAdmin(email, password)

	if err != nil {
		log.Fatalf("❌ No se pudo crear el admin inicial: %v", err)
	} else {
		log.Printf("✅ Se cargo exitosamente el admin inicial")
	}
}

