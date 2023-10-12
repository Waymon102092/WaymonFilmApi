package main

import (
	"Waymon_api/config"
	"Waymon_api/routes"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	port := viper.GetString("server.port")
	r := routes.InitRoutes()
	r.Run(port)
}
