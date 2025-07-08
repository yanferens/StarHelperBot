package handlers

import (
	tg "github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/yanferens/StarHelperBot/internal/service"
)

func InitMenuHandlers(bh *th.BotHandler, bot *tg.Bot) {
	// Handle /start command
	bh.HandleMessage(func(ctx *th.Context, msg tg.Message) error {
		text, buttons := service.MenuService()
		_, err := bot.SendMessage(ctx, tu.Message(
			tu.ID(msg.Chat.ID),
			text["Message"],
		).WithReplyMarkup(tu.InlineKeyboard(tu.InlineKeyboardRow(tu.InlineKeyboardButton("Чат").WithURL(buttons["Чат"]), tu.InlineKeyboardButton("Правила").WithURL(buttons["Правила чату"])), tu.InlineKeyboardRow(tu.InlineKeyboardButton("Магазин").WithCallbackData(buttons["Магазин"]), tu.InlineKeyboardButton("Команди").WithCallbackData(buttons["Команди"])))))
		if err != nil {
			return err
		}
		return nil

	}, th.CommandEqual("start"))
}
