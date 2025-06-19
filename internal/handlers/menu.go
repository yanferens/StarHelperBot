package handlers

import (
	tg "github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func InitMenuHandlers(bh *th.BotHandler, bot *tg.Bot) {
	// Handle /start command
	bh.HandleMessage(func(ctx *tg.Context, msg *tg.Message) error {

		
}, th.CommandEqual("start"))
	}