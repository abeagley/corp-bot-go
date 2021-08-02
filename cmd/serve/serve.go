package main

import (
	"corp_bot/lib"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := lib.Server{
		AppId:            os.Getenv("BOT_APP_ID"),
		FirebaseCredPath: os.Getenv("FIREBASE_CRED_FILE"),
		Token:            os.Getenv("BOT_TOKEN"),
	}
	s.StartBot()
}
