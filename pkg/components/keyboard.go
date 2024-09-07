package keyboard

import (
	

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/reply"
	"github.com/pyKis/test_bot/pkg/handler"
)

var demoReplyKeyboard *reply.ReplyKeyboard

func InitReplyKeyboard(b *bot.Bot) {
	demoReplyKeyboard = reply.New(
		b,
		reply.WithPrefix("reply_keyboard"),
		reply.IsSelective(),
		reply.IsOneTimeKeyboard(),
	).
		Button("Button", b, bot.MatchTypeExact, OnReplyKeyboardSelect).
		Row().
		Button("Cancel", b, bot.MatchTypeExact, OnReplyKeyboardSelect)
}


