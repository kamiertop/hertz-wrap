package main

import (
	"fmt"

	"hertz/internal/router"
	"hertz/pkg/config"
	"hertz/pkg/log"
)

func Init() error {
	// config
	if err := config.InitConfig(); err != nil {
		return err
	}
	// logger
	log.InitLogger()

	return nil
}

func main() {
	if err := Init(); err != nil {
		fmt.Printf("init server failed, err: %v\n", err)
	}
	engine := router.Init()

	_ = engine.Run()
}
