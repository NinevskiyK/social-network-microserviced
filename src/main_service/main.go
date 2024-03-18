package main

import (
	"log"
	controllers "main_service/controllers"
	db_utils "main_service/db_utils"
	middlewares "main_service/middlewares"

	gin "github.com/gin-gonic/gin"
)

func main() {
	err := db_utils.StartUpDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/user/register", controllers.Register)
	r.POST("/user/login", controllers.LoginUser)

	authed := r.Group("")
	authed.Use(middlewares.AuthMiddleware)
	authed.PUT("/user/update_me", controllers.UpdateMe)

	r.Run(":8081")
}
