package system

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func BotToken() string {
	err := godotenv.Load(".env")

	if err != nil{
		log.Print("No .env file found")
	}
	
	token := os.Getenv("BOT_TOKEN")


	return token


}