package handler

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
   k "github.com/pyKis/test_bot/pkg/bot"

)


func Handler(ctx context.Context, b *bot.Bot, update *models.Update)  {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}

func HandlerReplyKeyboard(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      update.Message.Chat.ID,
		Text:        "Select example command from reply keyboard:",
		ReplyMarkup: k.ReplyKeyboard,
	})
}

func OnReplyKeyboardSelect(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "You selected: " + string(update.Message.Text),
	})
}