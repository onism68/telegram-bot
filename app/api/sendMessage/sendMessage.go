package sendMessage

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
	"telegram-bot/library/telegram"
)

// Hello is a demonstration route handler for output "Hello World!".
func SendMessage(r *ghttp.Request) {
	if r.Method == http.MethodPost {
		id := r.GetInt64("id")
		msg := r.GetString("msg")
		telegram.SendMessage(
			tgbotapi.NewMessage(id, msg))
		r.Response.Writeln("Hello World!")
	}

}
