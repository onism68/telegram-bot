package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"golang.org/x/net/proxy"
	"net/http"
	"os"
	"sync"
)

type Message struct {
	messageChan chan interface{}
}

type Bot struct {
	*tgbotapi.BotAPI

	start bool
}

var Instance *Bot

func BotInit() (*Bot, error) {
	// 设置代理
	socksProxy := os.Getenv("socksProxy")
	tgToken := os.Getenv("tgToken")
	socks5, err := proxy.SOCKS5("tcp", socksProxy, nil, proxy.Direct)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{}
	transport.Dial = socks5.Dial
	// 初始化bot
	bot, err := tgbotapi.NewBotAPIWithClient(tgToken, &http.Client{Transport: transport})
	if err != nil {
		return nil, err
	}
	bot.Debug = false
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 10
	updatesChan := bot.GetUpdatesChan(u)
	// 维护一个全局的channel
	// 收到消息后进行命令判断“/”
	// 将命令与module中进行比对？
	// 符合，执行module中的相关操作（将bot示例也传过去？）

	// module的注册好像成了摆设？
	for _, moduleInfo := range Modules {
		moduleInfo.Instance.Init()
	}

	for _, moduleInfo := range Modules {
		moduleInfo.Instance.PostInit()
	}

	for _, moduleInfo := range Modules {
		moduleInfo.Instance.Serve(&Bot{
			bot, true,
		})
	}

	go updateMessage(&Bot{
		bot, true,
	}, updatesChan)

	Instance = &Bot{
		bot, true,
	}

	return &Bot{
		bot, true,
	}, nil
}

// Stop 停止所有服务
func (*Bot) Stop() {
	glog.Warning("stopping ...")
	wg := sync.WaitGroup{}
	for _, moduleInfo := range Modules {
		wg.Add(1)
		moduleInfo.Instance.Stop(Instance, &wg)
	}
	wg.Wait()
	glog.Info("stopped")
	Modules = make(map[string]ModuleInfo)
}

func updateMessage(bot *Bot, updatesChan tgbotapi.UpdatesChannel) {
	for update := range updatesChan {
		//glog.Debug(update)
		if update.Message == nil {
			continue
		}

		for _, moduleInfo := range Modules {

			// 注册到全局的module
			if moduleInfo.Id.Namespace() == GlobalModule {
				moduleInfo.Instance.Start(bot, update)
			}
			// 判断消息是否为command
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case moduleInfo.Id.Name():
					moduleInfo.Instance.Start(bot, update)
				}
			}
		}
	}
}

func SendMessage(msg tgbotapi.MessageConfig) {
	glog.Infof("[send] [%d]: %s", msg.ChatID, msg.Text)
	if _, err := Instance.Send(msg); err != nil {
		glog.Error(err)
	}
}
