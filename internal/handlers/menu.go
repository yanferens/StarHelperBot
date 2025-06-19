package handlers

import (
	tg "github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func InitMenuHandlers(bh *th.BotHandler, bot *tg.Bot) {
	// Handle /start command
	bh.HandleMessage(func(ctx *th.Context, msg tg.Message) error {
		_, err := bot.SendMessage(ctx, tu.Messagef(
			tu.ID(msg.Chat.ID),
			"Welcome to StarHelperBot! Use /help to see available commands.",
		))
		if err != nil {
			return err
		}
		return nil
		
}, th.CommandEqual("start"))
	}