package main

import (
	"praticeone/action"
	"praticeone/common"

	"github.com/gin-gonic/gin"
)

func main() {
	// Read json file.
	common.ParseJSONFile()

	router := gin.Default()

	router.GET("/role", action.GetAllRole)

	router.GET("/role/:id", action.GetRoleById)

	router.POST("/role", action.PostRole)

	router.PUT("/role/:id", action.PutRole)

	router.DELETE("/role/:id", action.DeleteRole)

	router.Run(":8080")
}
