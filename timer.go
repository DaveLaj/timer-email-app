package main

import (
	"time"
)

func TimerEmailLoop(cfg Config) {
	go func() {
		for {
			EmailOnTimer(cfg)
			time.Sleep(time.Duration(cfg.SEND_INTERVAL) * time.Second)
		}
	}()
}
