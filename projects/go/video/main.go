package main

import (
	"github.com/caovanhoang63/hiholive/video/cmd"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := godotenv.Load("./projects/go/user/.env.local"); err != nil {
		log.Fatal("Error loading .env file")
	}
	cmd.Execute()
}
