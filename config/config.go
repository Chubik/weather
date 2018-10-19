package config

import "github.com/spf13/viper"

//Config structure for the file configuration variables
type Config struct {
	WebServerAddr string
	FileName      string
}

//C is a global variable that contains all initialization params of APP
var C Config

//InitialConfig function loading file configuration and setting variables
func InitialConfig() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()

	C.WebServerAddr = v.GetString("addr")
	C.FileName = v.GetString("file")

	return err
}
