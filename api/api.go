package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//PingHandler simple ping/pong endpoint
func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

//NoRoute default answer page not found
func NoRoute(c *gin.Context) {
	c.String(http.StatusOK, "Route not found")
}
