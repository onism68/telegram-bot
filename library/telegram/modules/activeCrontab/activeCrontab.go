package activeCrontab

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *ActiveCrontab
var id = telegram.ActiveCron + ".AC"

func init() {
	glog.Debugf("开始注册: %s", id)
	telegram.RegisterModule(instance)
}

type ActiveCrontab struct{}

func (c *ActiveCrontab) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       telegram.ModuleId(id),
		Instance: instance,
	}
}

func (c *ActiveCrontab) Init() {
	//panic("implement me")
}

func (c *ActiveCrontab) PostInit() {
	//panic("implement me")
}

func (c *ActiveCrontab) Serve(bot *telegram.Bot) {
	//panic("implement me")
}

func (c *ActiveCrontab) Start(bot *telegram.Bot, update tgbotapi.Update) {
	//panic("implement me")
	glog.Info("准备激活crontab任务, 设置id")
	crontabChatId := update.Message.Chat.ID
	// crontab modules
	go func() {
		for _, moduleInfo := range telegram.Modules {
			if moduleInfo.Id.Namespace() == telegram.CronModule {
				// crontab modules 只应该依赖Update的chat id
				go moduleInfo.Instance.Start(bot,
					tgbotapi.Update{
						Message: &tgbotapi.Message{
							Chat: &tgbotapi.Chat{
								ID: crontabChatId}}})
			}
		}
	}()
	msg := tgbotapi.NewMessage(crontabChatId, "已激活 crontab 任务")
	_, err := bot.Send(msg)
	if err != nil {
		glog.Error(err)
	}
}

func (c *ActiveCrontab) Stop(bot *telegram.Bot, wg *sync.WaitGroup) {
	//panic("implement me")
}
