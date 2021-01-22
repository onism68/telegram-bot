package memorandum

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"strings"
	"sync"
	"telegram-bot/library/telegram"
	"time"
)

var instance *Memorandum
var id = "func.memorandum"

func init() {
	glog.Debugf("开始注册: %s", id)
	telegram.RegisterModule(instance)
}

type Memorandum struct{}

func (m *Memorandum) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       telegram.ModuleId(id),
		Instance: instance,
	}
}

func (m *Memorandum) Init() {
	//panic("implement me")
}

func (m *Memorandum) PostInit() {
	//panic("implement me")
}

func (m *Memorandum) Serve(bot *telegram.Bot) {
	//panic("implement me")
}

func (m *Memorandum) Start(bot *telegram.Bot, update tgbotapi.Update) {
	do(update)
}

func (m *Memorandum) Stop(bot *telegram.Bot, wg *sync.WaitGroup) {
	//panic("implement me")
}

func do(update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "例如: /memorandum 2020-01-20 20:20:20||请提醒我")
	args := update.Message.CommandArguments()
	argsTmpList := strings.Split(args, "||")
	if len(argsTmpList) == 2 {
		timeFlag := gtime.ParseTimeFromContent(argsTmpList[0])
		content := argsTmpList[1]
		glog.Info(timeFlag, content)
		flagDuration := gconv.Duration(timeFlag.Timestamp()-gtime.Timestamp()) * time.Second
		msgText := fmt.Sprintf("[timer] 将在 [%s] (%s后) 提醒你 [%s]", argsTmpList[0], flagDuration, content)
		glog.Info(msgText)
		msg.Text = msgText
		gtimer.AddOnce(flagDuration, func() {
			msg.Text = content
			telegram.SendMessage(msg)
		})
	}
	telegram.SendMessage(msg)
}
