package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	//_ "telegram-bot/boot"
	"telegram-bot/library/telegram"
	//_ "telegram-bot/router"
)

func main() {
	bot, err := telegram.Bot()
	if err != nil {
		panic(err)
	}
	glog.Info(bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10

	updatesChan := bot.GetUpdatesChan(u)

	for update := range updatesChan {
		if update.Message == nil {
			continue
		}
		glog.Infof("ChatId[%d] %s", update.Message.Chat.ID, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		send, err := bot.Send(msg)
		if err != nil {
			glog.Error(err)
		}
		glog.Info(send.Text)
	}

	//g.Server().Run()
}
