package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/net/proxy"
	"net/http"
	"os"
)

func Bot() (*tgbotapi.BotAPI, error) {
	socksProxy := os.Getenv("socksProxy")
	tgToken := os.Getenv("tgToken")
	socks5, err := proxy.SOCKS5("tcp", socksProxy, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{}
	transport.Dial = socks5.Dial
	bot, err := tgbotapi.NewBotAPIWithClient(tgToken, &http.Client{Transport: transport})
	if err != nil {
		return nil, err
	}
	bot.Debug = false
	return bot, nil
}
