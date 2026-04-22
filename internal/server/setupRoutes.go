package server

import (

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/handlers"
	"github.com/tu-usuario/blog-api/internal/middleware"
	"github.com/tu-usuario/blog-api/internal/models"
)

func SetupRoutes(frontHost string) *gin.Engine {
	r := gin.Default()

	setupCORS(r, frontHost)

	setupPublicRoutes(r)

	setupProtectedRoutes(r)

	setupAdminRoutes(r)

	return r
}


func setupCORS(r *gin.Engine, frontHost string) {
	// metodo que configura CORS para poder correr REACT sin problema de forma Local a la ves que el servidor
	
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontHost}, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func setupPublicRoutes(r *gin.Engine) {
	public := r.Group("/api/v1")
	{
		public.GET("/posts", handlers.RetreiveAllPosts)
	
		public.GET("/posts/:id", handlers.GetPostByID)
	
		public.GET("/posts/search", handlers.SearchPosts)
	
		public.POST("/register", handlers.RegisterUser)

		public.POST("/login", handlers.Login)
	}
}

func setupProtectedRoutes(r *gin.Engine) {
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthJWT())

	protected.POST("/posts", handlers.CreatePost)
}

func setupAdminRoutes(r *gin.Engine) {
	admin := r.Group("/api/v1/admin")
	admin.Use(middleware.AuthJWT())
	admin.Use(middleware.RequireRole(models.RoleAdmin))

	admin.GET("/users", handlers.GetUsers)

	admin.DELETE("/posts", handlers.DeleteAllPosts)

	admin.DELETE("/posts/:id", handlers.DeleteByID)

}
