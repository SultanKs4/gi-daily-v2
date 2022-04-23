package setting

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(rootDir()); err != nil {
		log.Printf("[INFO] %v", err)
	}
}

func Setup() {
	Environment.Cookie = os.Getenv("COOKIE")
	Environment.TelegramBotToken = os.Getenv("TG_BOT_TOKEN")
	Environment.TelegramOwnerId = os.Getenv("TG_OWNER_ID")
}
