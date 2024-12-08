package main

import (
	"github.com/caovanhoang63/hiholive/services/hls_mux/cmd"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()
}
