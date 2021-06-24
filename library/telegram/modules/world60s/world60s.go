package world60s

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *World60s
var id = telegram.CronModule + ".world60s"

func init() {
	glog.Debugf("开始注册: %s", id)
	telegram.RegisterModule(instance)
}

type World60s struct{}

func (w *World60s) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       telegram.ModuleId(id),
		Instance: instance,
	}
}

func (w *World60s) Init() {
	//panic("implement me")
}

func (w *World60s) PostInit() {
	//panic("implement me")
}

func (w *World60s) Serve(bot *telegram.Bot) {
	//panic("implement me")
}

func (w *World60s) Start(bot *telegram.Bot, update tgbotapi.Update) {
	glog.Info("crontab world60s")
	_, err := gcron.Add("0 0 9 * * *", func() {
		glog.Info("world60s crontab")
		do(bot, update)
	})
	if err != nil {
		glog.Error(err)
	}

}

func (w *World60s) Stop(bot *telegram.Bot, wg *sync.WaitGroup) {
	//panic("implement me")
}

func do(bot *telegram.Bot, update tgbotapi.Update) {
	url := "http://api.03c3.cn/zb/"
	upload := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(url))
	_, err := bot.Send(upload)
	if err != nil {
		glog.Errorf("bot.Send(upload) error %s", err.Error())
	}
}
