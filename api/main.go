package main

import (
	"errors"
	"log"

	"github.com/tu-usuario/blog-api/internal/auth"
	"github.com/tu-usuario/blog-api/internal/config"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/server"
)

func main() {

	cfg := config.Load()

	initDB(cfg.DataBaseURI)
	configAuth(cfg.AdminEmail, cfg.AdminPassword, cfg.JWTKey)

	setUpServer(cfg.Host, cfg.Port, cfg.FrontHost)
}

func initDB(uri string){
	err := db.InitDB(uri)

	if err!= nil{
		switch{
			case errors.Is(err, db.ErrFailedToConnectToDataBase):
				log.Fatal("❌ No se pudo conectar a MongoDB")
			case errors.Is(err, db.ErrFailedToCreateIndexis):
				log.Fatal("❌ Failed making emails indexis")
			default:
				log.Fatalf("❌ Unknown error: %v /n", err)
		}
	}

	log.Println("✅ Conectado a MongoDB exitosamente")
}


func configAuth(email string, password string, jwtKey string) {

	auth.LoadJWT(jwtKey)

	err := auth.SeedAdmin(email, password)

	if err != nil {
		log.Fatalf("❌ No se pudo crear el admin inicial: %v", err)
	} else {
		log.Printf("✅ Se cargo exitosamente el admin inicial")
	}
}


func setUpServer(host string, port string, frontHost string) {
	r := server.SetupRoutes(frontHost)

	log.Println("✅ CORS allows origin: " + frontHost)

	err := r.Run(host + ":" + port)
    if err != nil {
        log.Fatalf("❌ El servidor se detuvo de forma inesperada: %v", err)
    }
}