package main

import (
	"log"

	"github.com/joho/godotenv"
	appwire "github.com/yodzafar/food-marketpalce/user-service/wire"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}
	app, err := appwire.InitApp()
	if err != nil {
		log.Fatalf("init app: %v", err)
	}
	app.Run()
}
