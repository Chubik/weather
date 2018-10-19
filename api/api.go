package api

import (
	"fmt"
	"net/http"
	"strings"
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
	pc := q["str"][0]
	for _, city := range data.Cities {
		ct := strings.Split(city.Name, ",")
		if strings.ToLower(pc) == strings.ToLower(ct[0]) {
			c.String(http.StatusOK, fmt.Sprintf("%v\n", city))
			return
		}
	}
	c.String(http.StatusOK, "Cities not found")
}
