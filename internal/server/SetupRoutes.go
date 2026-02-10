package server

import (
	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/handlers"
	"github.com/tu-usuario/blog-api/internal/middleware"
	"github.com/tu-usuario/blog-api/internal/models"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	setupPublicRoutes(r)

	setupProtectedRoutes(r)

	setupAdminRoutes(r)

	return r
}

func setupPublicRoutes(r *gin.Engine) {
	public := r.Group("/api/v1")
	{
		public.GET("/posts", handlers.RetreiveAllPosts)
	
		public.GET("/posts/:id", handlers.GetPostByID)
	
		public.GET("/posts/search", handlers.SearchPosts)
	
		public.POST("/regiser", handlers.RegisterUser)

		public.POST("/create-admin", handlers.CreateAdmin)

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
