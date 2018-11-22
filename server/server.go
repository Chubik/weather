package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"weather/api"
	"weather/config"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

type WServer http.Server

var (
	WeatherServer WServer
)

const (
	healthCheckEndPoint = "/ping"
)

//Init server initialization function
func (ws *WServer) Init() {

	//Change on production deploy
	// gin.SetMode(gin.ReleaseMode)

	ws.Addr = config.C.WebServerAddr

	// r := gin.New()
	r := gin.Default()

	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	s := &http.Server{
		Addr:           ws.Addr,
		Handler:        r,
		ReadTimeout:    15 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//setup default page not found
	r.NoRoute(api.NoRoute)
	//WEBServer Health Check route
	r.GET(healthCheckEndPoint, api.PingHandler)

	//Routes
	r.GET("/getcities", api.GetCity)

	fmt.Printf("[INFO] Server starting: %v\n", ws.Addr)
	//Start WEB server
	log.Fatal(s.ListenAndServe())

}
