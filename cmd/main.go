package main

import (
	"fmt"
	"github.com/Klasos93/4life/config"
	"github.com/Klasos93/4life/internal/logger"
)

func main() {
	cfg, err := config.MustLoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config loaded: %#v\n", cfg)
	logger.InitLogger()
	logger.Log.Info("Logger initialized")
	//TODO: create connect db
	//TODO: create Makefile

}
