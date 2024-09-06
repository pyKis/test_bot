package system

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func BotToken () string {
	err := godotenv.Load(".env")

	if err != nil{
		log.Print("No .env file found")
	}
	
	return os.Getenv("BOT_TOKEN")
	

}