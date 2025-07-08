package main

import (
	"context"
	"log"
	"net/http"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	handlers "github.com/yanferens/StarHelperBot/internal/handler"
)

func main() {
	ctx := context.Background()
	token := TOKEN
	if token == "" {
		log.Fatal("TOKEN env var required")
	}

	bot, err := telego.NewBot(token, telego.WithDefaultDebugLogger())
	// Set webhook
	err = bot.SetWebhook(ctx, &telego.SetWebhookParams{
		URL:         "https://4542-93-127-124-136.ngrok-free.app/bot",
		SecretToken: bot.SecretToken(),
	})
	if err != nil {
		log.Fatalf("set webhook error: %v", err)
	}

	mux := http.NewServeMux()

	updates, err := bot.UpdatesViaWebhook(ctx, telego.WebhookHTTPServeMux(mux, "/bot", bot.SecretToken()))
	if err != nil {
		log.Fatalf("updates webhook error: %v", err)
	}

	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("http listen error: %v", err)
		}
	}()

	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)
	handlers.InitMenuHandlers(bh, bot)

	// Stop handling updates
	defer func() { _ = bh.Stop() }()
	//InitHandlers(bh, bot)

	_ = bh.Start()
}
