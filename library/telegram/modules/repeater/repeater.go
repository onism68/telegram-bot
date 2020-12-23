package repeater

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *repeater
var id = "cmd.repeater"

func init() {
	glog.Debugf("开始注册: %s", id)
	telegram.RegisterModule(instance)
}

type repeater struct{}

func (*repeater) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       telegram.ModuleId(id),
		Instance: instance,
	}
}

func (*repeater) Init() {
	//panic("implement me")
}

func (*repeater) PostInit() {
	//panic("implement me")
}

func (*repeater) Serve(bot *telegram.Bot) {
	//panic("implement me")
}

func (*repeater) Start(bot *telegram.Bot, update tgbotapi.Update) {
	// 判断收到的信息大于命令长度, 避免空命令
	if len(update.Message.Text) > update.Message.Entities[0].Length+1 {
		msgTmp := update.Message.Text[update.Message.Entities[0].Length+1:]
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgTmp)
		if _, err := bot.Send(msg); err != nil {
			glog.Error(err)
		}
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "使用方法: /repeater 复读文字")
		if _, err := bot.Send(msg); err != nil {
			glog.Error(err)
		}
	}
}

func (*repeater) Stop(bot *telegram.Bot, wg *sync.WaitGroup) {
	//panic("implement me")
}
