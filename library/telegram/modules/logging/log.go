package logging

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *tgLog

func init() {
	//instance := new(tgLog)
	glog.Debug("开始注册: global.logging")
	telegram.RegisterModule(instance)
}

type tgLog struct{}

func (*tgLog) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       "global.logging",
		Instance: instance,
	}
}

func (*tgLog) Init() {
	glog.Info("[logging] init ")
}

func (*tgLog) PostInit() {
	glog.Info("[logging] postInit ")
}

func (*tgLog) Serve(bot *telegram.Bot) {
	glog.Info("[logging] serve ")
}

func (*tgLog) Start(bot *telegram.Bot, update tgbotapi.Update) {
	go glog.Infof("[logging] [%s]-[%d]: %s", update.Message.Chat.UserName, update.Message.Chat.ID, update.Message.Text)
}

func (*tgLog) Stop(bot *telegram.Bot, wg *sync.WaitGroup) {
	defer wg.Done()
	glog.Info("[logging] stopping ")
}
