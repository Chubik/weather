package main

import (
	"log"
	"weather/config"
	"weather/data"
	"weather/server"
)

func main() {

	// rand.Seed(time.Now().UnixNano())

	// core.LNQueue = &core.LNStack{}
	// core.Wallet = &core.JackPot{}
	err := config.InitialConfig()
	if err != nil {
		log.Fatal("[ERROR config init]", err)
	}
	err = data.Init(config.C.FileName)
	if err != nil {
		log.Fatal("[ERROR] data init: ", err)
	}
	// setup.AllDataSetInit(core.LNQueue, core.Wallet, setup.C.TimerSec)
	server.WeatherServer.Init()
}
