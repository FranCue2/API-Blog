package middleware

import (

	"github.com/gin-gonic/gin"
	"github.com/tu-usuario/blog-api/internal/auth"
	"github.com/tu-usuario/blog-api/internal/constants"
	"github.com/tu-usuario/blog-api/internal/models"
)


func AuthJWT() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")

		claims, err := auth.ValidateToken(tokenString)

		if err != nil{
			c.AbortWithStatusJSON(401, gin.H{"error":"Sesion vencida", "err": err.Error()})
			return 
		}

		c.Set(constants.UserID, claims.Subject)
		c.Set(constants.UserRole, claims.Role)
		c.Set(constants.UserEmail, claims.Email)

		c.Next()
	}
}

func RequireRole(requiredRole models.Role) gin.HandlerFunc{
	return func(c *gin.Context){
		val, exists := c.Get(constants.UserRole)
		
		if !exists{
			c.AbortWithStatusJSON(401, gin.H{"error":"No se pudo validar el rol del usuario"})
			return 
		}

		userRole, ok := val.(models.Role)
		
		if !ok {
			c.AbortWithStatusJSON(500, gin.H{"error":"formato de rol invalido"})
			return
		}
		
		if (models.Role(userRole)!= requiredRole){
			c.AbortWithStatusJSON(403, gin.H{"error":"No tenes permisos suficientes para esta accion"})
			return
		}

		c.Next()
	}
}