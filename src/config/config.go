package config

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	coolify "coolifymanager/src/coolity"
	_ "github.com/joho/godotenv/autoload"
)

// helper: env with default value
func env(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

var (
	Coolify *coolify.Client

	// SAME environment variables (unchanged names)
	ApiUrl     = env("API_URL", "https://app.coolify.io")
	ApiToken   = env("API_TOKEN", "1|k4tLaFHqb1jz8mNitFYDmKHQz4aqcQY9nM1ga7VWd9b28be0")
	Token      = env("TOKEN", "8104460140:AAEnI5F2oBkRSKMPgHh6L5O3s6D_5-ap8XA")
	Port       = env("PORT", "8000")
	WebhookUrl = env("WEBHOOK_URL", "https://protech.eu.org/")
	devList    = env("DEV_IDS", "6035523795") // comma-separated
	devIDs     []int64
)

func Init() error {
	// Required checks
	if ApiUrl == "" {
		return errors.New("API_URL must be set")
	}
	if ApiToken == "" {
		return errors.New("API_TOKEN must be set")
	}
	if Token == "" {
		return errors.New("TOKEN must be set")
	}

	// Init Coolify client
	Coolify = &coolify.Client{
		BaseURL: ApiUrl,
		Token:   ApiToken,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// Parse DEV_IDS
	devIDs = devIDs[:0] // reset slice (safe on reload)
	for _, idStr := range strings.Split(devList, ",") {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			log.Printf("Invalid DEV_ID: %s", idStr)
			continue
		}
		devIDs = append(devIDs, id)
	}

	log.Printf("Config loaded | API_URL=%s | PORT=%s | DEV_IDS=%v",
		ApiUrl, Port, devIDs)

	return nil
}

// IsDev checks if a given Telegram user ID is a developer
func IsDev(userID int64) bool {
	for _, id := range devIDs {
		if id == userID {
			return true
		}
	}
	return false
}
