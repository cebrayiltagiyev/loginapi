package main

import (
	"LoginRegister/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main(){

	router:=gin.Default()

	loginController := controller.LoginController{}

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session_id", store))

	router.POST("/login" , loginController.LoginByUsername)

	router.Run(":8000")



}