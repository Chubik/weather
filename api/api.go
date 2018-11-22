package api

import (
	"net/http"
	"strings"
	"weather/config"
	"weather/data"

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

//GetCity function that finding part of cities and return a slice of cities
func GetCity(c *gin.Context) {
	q := c.Request.URL.Query()
	pc := strings.ToLower(q["str"][0])
	if len(pc) < config.C.Cap {
		c.String(http.StatusOK, "Short name of City. Should be more > "+string(config.C.Cap))
		return
	}
	resp := data.Search(pc)
	if len(resp) > 0 {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.String(http.StatusOK, "Cities not found")
}
