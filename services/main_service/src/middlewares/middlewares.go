package middlewares

import (
	"main_service/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "got no token"})
		return
	}

	id, err := utils.GetIdFromJWT(token)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		return
	}

	c.Set("id", id)
	c.Next()
}
