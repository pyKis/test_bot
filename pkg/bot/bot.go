package bot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	
	"github.com/pyKis/test_bot/pkg/system"
)



func BotStart(){
	token := system.BotToken()
	
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler.Handler),
	}

	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}