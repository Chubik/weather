package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
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
	gin.SetMode(gin.ReleaseMode)

	ws.Addr = GetPort()
	r := gin.New()
	//r := gin.Default()

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

	r.LoadHTMLFiles("index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	//WEBServer Health Check route
	r.GET(healthCheckEndPoint, api.PingHandler)

	//Routes
	r.GET("/getcities", api.GetCity)

	//WebSockets setup
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	fmt.Printf("[INFO] Server starting: %v\n", ws.Addr)
	//Start WEB server
	log.Fatal(s.ListenAndServe())

}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = config.C.WebServerAddr
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
