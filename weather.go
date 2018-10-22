package main

import (
	"log"
	"weather/config"
	"weather/data"
	"weather/server"
)

func main() {

	err := config.InitialConfig()
	if err != nil {
		log.Fatal("[ERROR config init]", err)
	}
	err = data.Init(config.C.FileName)
	if err != nil {
		log.Fatal("[ERROR] data init: ", err)
	}
	server.WeatherServer.Init()
}
