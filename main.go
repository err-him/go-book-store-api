package main

import (
	"book-store-api/api/utils"
	"book-store-api/config"
	"log"
	"os"

	"github.com/err-him/gonf"
)

type PortConfig struct {
	Port string
}

func main() {

	//Get port from command line interface
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "prod"
	}
	cfg := PortConfig{}
	err := gonf.GetConfig(utils.GetEnvFile(env), &cfg)
	if err != nil {
		log.Fatal("environment can not be loaded at this moment, Please try after some time", err)
	}
	app := &config.App{}
	app.Intialize()
	app.Run(cfg.Port)

}
