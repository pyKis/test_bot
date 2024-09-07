package bot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/ui/keyboard/reply"

	"github.com/pyKis/test_bot/pkg/handler"
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
	InitReplyKeyboard(b)
	b.Start(ctx)
	
}

var ReplyKeyboard *reply.ReplyKeyboard

func InitReplyKeyboard(b *bot.Bot) {
	ReplyKeyboard = reply.New(
		b,
		reply.WithPrefix("reply_keyboard"),
		reply.IsSelective(),
		reply.IsOneTimeKeyboard(),
	).
		Button("Button", b, bot.MatchTypeExact, handler.OnReplyKeyboardSelect).
		Row().
		Button("Cancel", b, bot.MatchTypeExact, handler.OnReplyKeyboardSelect)
}
