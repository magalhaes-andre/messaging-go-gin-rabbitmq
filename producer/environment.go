package main

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("No .env file was found: %v", err)
	}
}
