package hello

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/net/ghttp"
	"telegram-bot/library/telegram"
)

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
	telegram.SendMessage(
		tgbotapi.NewMessage(404176520, "1"))
	r.Response.Writeln("Hello World!")
}
