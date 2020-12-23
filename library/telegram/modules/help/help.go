package repeater

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *help
var id = "cmd.help"

func init() {
	glog.Debugf("开始注册: %s", id)
	telegram.RegisterModule(instance)
}

type help struct{}

func (*help) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       telegram.ModuleId(id),
		Instance: instance,
	}
}

func (*help) Init() {
	//panic("implement me")
}

func (*help) PostInit() {
	//panic("implement me")
}

func (*help) Serve(bot *telegram.Bot) {
	//panic("implement me")
}

func (*help) Start(bot *telegram.Bot, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"/help 帮助\n"+
			"/repeater hello 复读: hello")
	if _, err := bot.Send(msg); err != nil {
		glog.Error(err)
	}
}

func (*help) Stop(bot *telegram.Bot, wg *sync.WaitGroup) {
	//panic("implement me")
}
