package config

import "github.com/spf13/viper"

//Config structure for the file configuration variables
type Config struct {
	WebServerAddr string
	WebSocketAddr string
	FileName      string
	Cap           int
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
	C.WebSocketAddr = v.GetString("ws_addr")
	C.FileName = v.GetString("file")
	C.Cap = v.GetInt("cap")

	return err
}
