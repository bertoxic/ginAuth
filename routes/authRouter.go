package routes

import (
	 "github.com/bertoxic/ginAuth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine){
    incomingRoutes.POST("users/signup", controllers.SignUp())
    incomingRoutes.POST("users/login", controllers.Login())
}