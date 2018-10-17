package main

import (
	"log"
	"weather/config"
	"weather/server"
)

func main() {

	// rand.Seed(time.Now().UnixNano())

	// core.LNQueue = &core.LNStack{}
	// core.Wallet = &core.JackPot{}

	err := config.InitialConfig()
	if err != nil {
		log.Fatal(err)
	}
	// setup.AllDataSetInit(core.LNQueue, core.Wallet, setup.C.TimerSec)
	server.WeatherServer.Init()
}
