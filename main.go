package main

import (
	"AImBroke/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can't load .env file", err)
		return
	}
	loadedConfig := config.Load()

	fmt.Println("Loaded config: ", loadedConfig)
}
