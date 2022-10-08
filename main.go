package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	api := gin.New()
	api.GET("/", handle)

	err := api.Run(":8080")

	if err != nil {
		panic(err)
	}

}

func handle(c *gin.Context) {
	c.JSON(http.StatusOK, c.Request.Host)
}
