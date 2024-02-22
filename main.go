package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"go.uber.org/zap"

	"hertz/internal/router"
	"hertz/pkg/config"
	"hertz/pkg/log"
)

var pprof = flag.Bool("pprof", false, "start pprof")

func Init() error {
	// config
	if err := config.InitConfig(); err != nil {
		return err
	}
	// logger
	log.InitLogger()
	if *pprof {
		go func() {
			if err := http.ListenAndServe(":6060", nil); err != nil {
				log.Error("start pprof failed", zap.Error(err))
				os.Exit(1)
			}
		}()
	}

	return nil
}

func main() {
	if err := Init(); err != nil {
		fmt.Printf("init server failed, err: %v\n", err)
	}
	engine := router.Init()

	_ = engine.Run()
}
