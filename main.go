package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	api := gin.Default()
	api.GET("/", handle)
	api.GET("/hostname", hostname)
	api.GET("/addr", localaddr)

	err := api.Run(":80")

	if err != nil {
		panic(err)
	}

}

func handle(c *gin.Context) {
	c.JSON(http.StatusOK, c.Request.Host)
}

func hostname(c *gin.Context) {

	hostname, err := os.Hostname()

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, hostname)

}

func localaddr(c *gin.Context) {

	addr := c.Request.Context().Value(http.LocalAddrContextKey)
	host, port, err := net.SplitHostPort(fmt.Sprintf("%v", addr))

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, host+":"+port)

}
