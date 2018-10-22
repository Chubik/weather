package api

import (
	"encoding/json"
	"log"
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
	var resp []string
	q := c.Request.URL.Query()
	pc := strings.ToLower(q["str"][0])
	if len(pc) < 4 {
		c.String(http.StatusOK, "Short name of City. Should be more > 3")
		return
	}
	for _, city := range data.Cities {
		ct := strings.Split(city.Name, ",")

		if strings.Contains(strings.ToLower(ct[0]), pc) {
			// if strings.ToLower(pc) == strings.ToLower(ct[0]) {
			ret, err := json.Marshal(city)
			if err != nil {
				log.Println("Error parsing city data", err)
			}
			resp = append(resp, string(ret))
		}
	}
	if len(resp) > 0 {
		c.JSON(http.StatusOK, resp)
		return
	}
	c.String(http.StatusOK, "Cities not found")
}
