package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func main() {
	ctx := context.Background()
	

	// Note: Please keep in mind that default logger may expose sensitive information, use in development only
	bot, err := telego.NewBot(TOKEN, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
	// Get updates channel
	updates, _ := bot.UpdatesViaLongPolling(ctx, nil)

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	// Stop handling updates
	defer func() { _ = bh.Stop() }()
	//InitHandlers(bh, bot)

	_ = bh.Start()
}
