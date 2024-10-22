package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"hiholive/projects/go/hls_mux/cmd"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()
}
