package main

import (
	"log"

	"github.com/yodzafar/food-market/api-gateway/config"
	"github.com/yodzafar/food-market/api-gateway/internal/app"
)

func main() {
	cfg := config.Load()
	a, err := app.New(cfg)
	if err != nil {
		log.Fatalf("init: %v", err)
	}
	a.Run()
}
