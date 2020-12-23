package hello

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"telegram-bot/library/telegram"
)

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
	_, err := telegram.Instance.Send(tgbotapi.NewMessage(404176520, "1"))
	if err != nil {
		glog.Error(err)
	}
	r.Response.Writeln("Hello World!")
}
