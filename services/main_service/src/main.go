package main

import (
	"fmt"
	"log"
	controllers "main_service/controllers"
	db_utils "main_service/db_utils"
	middlewares "main_service/middlewares"
	"main_service/post_service"
	"os"

	gin "github.com/gin-gonic/gin"
)

func main() {
	err := db_utils.StartUpDB()
	if err != nil {
		log.Fatal(err)
	}
	err = post_service.Connect(fmt.Sprintf("%s:%s", os.Getenv("POST_SERVICE_HOST"), os.Getenv("POST_SERVICE_PORT")))
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.POST("/user/register", controllers.Register)
	r.POST("/user/login", controllers.LoginUser)

	authed := r.Group("")
	authed.Use(middlewares.AuthMiddleware)
	authed.PUT("/user/update_me", controllers.UpdateMe)
	authed.POST("/post/create", controllers.CreatePost)
	authed.PUT("/post/update/:post_id", controllers.UpdatePost)
	authed.DELETE("/post/delete/:post_id", controllers.DeletePost)
	authed.GET("/post/get/:post_id", controllers.GetPost)
	authed.GET("/wall", controllers.GetWall)
	r.Run(":8081")
}
