package app

import (
	"github.com/gin-gonic/gin"
	"os"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}
func StartApp() {
	mapUrls()
	port := os.Getenv("PORT")
	if err := router.Run(port); err != nil {
		panic(err)
	}
}
