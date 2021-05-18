package world60s

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
	"io/ioutil"
	"net/http"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *World60s
var id = "cron.world60s"

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
	_, err := gcron.Add("0 0 8 * * *", func() {
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
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	url := "http://api.03c3.cn/zb/"
	resp, err := http.Get(url)
	if err != nil {
		msg.Text = err.Error()
		telegram.SendMessage(msg)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		msg.Text = err.Error()
		telegram.SendMessage(msg)
	}
	bytes := tgbotapi.FileBytes{Name: "image.jpg", Bytes: b}
	upload := tgbotapi.NewPhoto(update.Message.Chat.ID, bytes)
	_, err = bot.Send(upload)
	if err != nil {
		msg.Text = err.Error()
		telegram.SendMessage(msg)
	}
}
