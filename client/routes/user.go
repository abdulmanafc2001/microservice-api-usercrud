package routes

import (
	"github.com/abdulmanafc2001/microservice-api-usercrud/client/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/create", controllers.CreateUser)
	router.PUT("/update/:id", controllers.UpdateUser)
	router.DELETE("/delete/:id", controllers.DeleteUser)
	router.GET("/", controllers.GetAllUsers)
}
