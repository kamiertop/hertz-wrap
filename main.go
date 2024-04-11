package main

import (
	"flag"
	"fmt"
	"os"

	"hertz/internal/router"
	"hertz/pkg/config"
	"hertz/pkg/log"
)

func Init() error {
	flag.Parse()
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
		os.Exit(1)
	}
	engine := router.Init()

	_ = engine.Run()
}
