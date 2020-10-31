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

	port := os.Getenv("PORT")
	errChan := make(chan error, 1)

	go func() {
		mapUrls()
		err := router.Run(port)
		errChan <- err
	}()
	if err := <-errChan; err != nil {
		panic(err)
	}
}
