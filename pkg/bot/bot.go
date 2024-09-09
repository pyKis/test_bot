package bot

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/go-telegram/ui/keyboard/reply"
	"github.com/pyKis/test_bot/pkg/system"
)


func BotStart(){
	token := system.BotToken()
	
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	
	opts := []bot.Option{
		bot.WithDefaultHandler(HandlerReplyKeyboard),
		bot.WithDefaultHandler(OnReplyKeyboardSelect),
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
		Button("Поделиться контактом", b, bot.MatchTypeExact, OnReplyKeyboardSelect).
		Row().
		Button("Сгенерировать ссылку", b, bot.MatchTypeExact, OnReplyKeyboardSelect)
}

func Handler(ctx context.Context, b *bot.Bot, update *models.Update)  {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

func HandlerReplyKeyboard(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Нажмите кнопку:",
		ReplyMarkup: ReplyKeyboard,
	})
}

func OnReplyKeyboardSelect(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Вы выбрали " + string(update.Message.Text),
	})
}