package config

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"coolifymanager/src/coolity"
	_ "github.com/joho/godotenv/autoload"
)

var (
	Coolify    *coolify.Client
	ApiUrl     = os.Getenv("API_URL", "https://app.coolify.io")
	ApiToken   = os.Getenv("API_TOKEN", "1|k4tLaFHqb1jz8mNitFYDmKHQz4aqcQY9nM1ga7VWd9b28be0")
	Token      = os.Getenv("TOKEN", "8104460140:AAEnI5F2oBkRSKMPgHh6L5O3s6D_5-ap8XA")
	Port       = os.Getenv("PORT", "8000")
	WebhookUrl = os.Getenv("WEBHOOK_URL", "https://protech.eu.org/")
	devList    = os.Getenv("DEV_IDS", "6035523795") // comma-separated
	devIDs     []int64                // parsed slice
)

func Init() error {
	if ApiUrl == "" || ApiToken == "" {
		return errors.New("API_URL and API_TOKEN must be set")
	}

	Coolify = &coolify.Client{
		BaseURL: ApiUrl,
		Token:   ApiToken,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// Parse DEV_IDS
	for _, idStr := range strings.Split(devList, ",") {
		idStr = strings.TrimSpace(idStr)
		if idStr == "" {
			continue
		}
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err == nil {
			devIDs = append(devIDs, id)
		} else {
			log.Printf("Dev ID is not an integer: %s", idStr)
		}
	}

	return nil
}

// IsDev checks if a given Telegram user ID is in the dev list
func IsDev(userID int64) bool {
	for _, id := range devIDs {
		if id == userID {
			return true
		}
	}
	return false
}
