package main

import (
	"github.com/abdulmanafc2001/microservice-api-usercrud/client/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.UserRoutes(router)
	router.Run(":4000")
}
