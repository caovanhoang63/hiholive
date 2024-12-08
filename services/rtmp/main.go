package main

import (
	"github.com/caovanhoang63/hiholive/services/rtmp/cmd"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()

}
