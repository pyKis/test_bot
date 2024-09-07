package main

import (
	"log"

	"github.com/pyKis/test_bot/pkg/system"
)



func main(){
	token:= system.BotToken()


	log.Printf("BOT_TOKEN: %s", token)
}