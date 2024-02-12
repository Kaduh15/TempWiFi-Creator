package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SSID      string
	USERNAME  string
	PASSWORD  string
	IP_ROUTER string
)

func InitDotEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
		panic(err)
	}

	SSID = os.Getenv("NAME_NETWORK")
	if SSID == "" {
		log.Fatal("NAME_NETWORK must be set")
		panic("NAME_NETWORK must be set")
	}

	USERNAME = os.Getenv("USER_ROUTER")
	if USERNAME == "" {
		log.Fatal("USER_ROUTER must be set")
		panic("USER_ROUTER must be set")
	}

	PASSWORD = os.Getenv("PASS_ROUTER")
	if PASSWORD == "" {
		log.Fatal("PASS_ROUTER must be set")
		panic("PASS_ROUTER must be set")
	}

	IP_ROUTER = os.Getenv("IP_ROUTER")
	if IP_ROUTER == "" {
		log.Fatal("IP_ROUTER must be set")
		panic("IP_ROUTER must be set")
	}
}
