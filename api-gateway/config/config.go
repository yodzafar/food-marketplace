package config

import "os"

type Config struct {
	HTTP     HTTPConfig
	Services ServicesConfig
}

type HTTPConfig struct {
	Port string
}

type ServicesConfig struct {
	UserAddr       string
	RestaurantAddr string
	MenuOrder      string
	OrderAddr      string
	ReviewAddr     string
}

func Load() *Config {
	return &Config{
		HTTP: HTTPConfig{
			Port: getEnv("PORT", "8080"),
		},
		Services: ServicesConfig{
			UserAddr:       getEnv("USER_ADDR", "localhost:50051"),
			RestaurantAddr: getEnv("RESTAURANT_ADDR", "localhost:50052"),
			MenuOrder:      getEnv("MENU_ORDER_ADDR", "localhost:50053"),
			OrderAddr:      getEnv("ORDER_ADDR", "localhost:50054"),
			ReviewAddr:     getEnv("REVIEW_ADDR", "localhost:50055"),
		},
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
