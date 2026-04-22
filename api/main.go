package main

import (
	"errors"
	"log/slog"
	"os"

	"github.com/tu-usuario/blog-api/internal/auth"
	"github.com/tu-usuario/blog-api/internal/config"
	db "github.com/tu-usuario/blog-api/internal/database"
	"github.com/tu-usuario/blog-api/internal/server"
)

func fatalErrorLog(loguer *slog.Logger,msg string, err error) {
	loguer.Error(msg, slog.String("error", err.Error()))
	os.Exit(1)
}

func main() {
	setupLogger()

	cfg := loadConfig()

	initDB(cfg.DataBaseURI)
	configAuth(cfg.AdminEmail, cfg.AdminPassword, cfg.JWTKey)

	setUpServer(cfg.Host, cfg.Port, cfg.FrontHost)
}

func setupLogger() {
	var logger *slog.Logger

	if os.Getenv("APP_ENV")=="production" {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	slog.SetDefault(logger)
}

func loadConfig() *config.AppConfig{
	cfg, err := config.Load()

	logConfig := slog.Default().With(slog.String("component", "config"))

	if err != nil {
		fatalErrorLog(logConfig, "Fallo al leer archivo .env", err)
	}else if(!cfg.IsProduction){
		logConfig.Info("Se cargo correctamente archivo .env")
	}else {
		logConfig.Info("Estamos en Modo Produccion")
	}

	return cfg
}

func initDB(uri string){
	logDB := slog.Default().With(slog.String("component", "database"))

	err := db.InitDB(uri)

	if err!= nil{
		switch{
			case errors.Is(err, db.ErrFailedToConnectToDataBase):
				fatalErrorLog(logDB, "No se pudo conectar con la Base de Datos", err)
			case errors.Is(err, db.ErrFailedToCreateIndexis):
				fatalErrorLog(logDB, "No se pudo poner como indices a los emails", err)
			default:
				fatalErrorLog(logDB, "Fallo desconocido", err)
		}
	}

	logDB.Info("Conectado a la Base de Datos Exitosamente")
}

func configAuth(email string, password string, jwtKey string) {
	logAuth := slog.Default().With(slog.String("component", "server"))


	auth.LoadJWT(jwtKey)

	err := auth.SeedAdmin(email, password)

	if err != nil {
		fatalErrorLog(logAuth,"No se pudo crear el admin inicial", err)
	} else {
		logAuth.Info("Se cargo exitosamente el admin inicial")
	}
}


func setUpServer(host string, port string, frontHost string) {
	logServer := slog.Default().With(slog.String("component", "server"))

	r := server.SetupRoutes(frontHost)

	logServer.Info("Configurado CORS", slog.String("host permitido",frontHost))

	err := r.Run(host + ":" + port)
    if err != nil {
        fatalErrorLog(logServer, "El servidor se detuvo de forma inesperada", err)
    }
}