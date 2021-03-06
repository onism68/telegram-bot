package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"net/url"
	"os"
	"sync"
	"telegram-bot/library/redisTool"
	"telegram-bot/library/telegram/types"
	"time"
)

type Message struct {
	messageChan chan interface{}
}

type Bot struct {
	*tgbotapi.BotAPI

	start bool
}

var instance *Bot

func BotInit() (*Bot, error) {
	// 设置代理
	socksProxy := os.Getenv("socksProxy")
	tgToken := os.Getenv("tgToken")
	var bot *tgbotapi.BotAPI
	var err error
	if socksProxy != "" {
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(socksProxy)
		}
		transport := &http.Transport{
			Proxy: proxy,
		}
		// 初始化bot
		bot, err = tgbotapi.NewBotAPIWithClient(tgToken, tgbotapi.APIEndpoint, &http.Client{Transport: transport})
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		bot, err = tgbotapi.NewBotAPI(tgToken)
		if err != nil {
			return nil, err
		}
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

	// 获取新消息
	go updateMessage(&Bot{
		bot, true,
	}, updatesChan)

	// 根据消息订阅发送新消息
	go chanSendMessage()

	instance = &Bot{
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
		moduleInfo.Instance.Stop(instance, &wg)
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
				go moduleInfo.Instance.Start(bot, update)
			}
			// 判断消息是否为command
			if update.Message.IsCommand() {
				switch update.Message.Command() {
				case moduleInfo.Id.Name():
					go moduleInfo.Instance.Start(bot, update)
				}
			}
		}
	}
}

func SendMessage(msg tgbotapi.MessageConfig) {
	glog.Infof("[send] [%d]: %s", msg.ChatID, msg.Text)
	if _, err := instance.Send(msg); err != nil {
		glog.Error(err)
	}
}

func chanSendMessage() {
	recMsgChan := make(chan types.TgMsg, 10)
	// 从channel中读取消息
	go func(recMsgChan chan types.TgMsg) {
		for chanMsg := range recMsgChan {
			msg := tgbotapi.NewMessage(chanMsg.ChatId, "")
			if chanMsg.Type == "" || chanMsg.Type == "text" {
				msgRunes := gconv.Runes(chanMsg.Message)
				if len(msgRunes) >= 4096 {
					msg.Text = string(msgRunes[:4096])
					SendMessage(msg)
					time.Sleep(100 * time.Millisecond)
					msg.Text = string(msgRunes[4096:])
					SendMessage(msg)
				} else {
					msg.Text = chanMsg.Message
					SendMessage(msg)
				}
			} else if chanMsg.Type == "img" {
				var mediaGroups []tgbotapi.MediaGroupConfig
				mediaGroup := tgbotapi.NewMediaGroup(chanMsg.ChatId, []interface{}{})
				for index, item := range chanMsg.ImgList {
					//resp, err := http.Get(item)
					//if err != nil {
					//	msg.Text = err.Error()
					//	SendMessage(msg)
					//	glog.Errorf("获取图片内容出错!", err)
					//}
					//defer resp.Body.Close()
					//b, err := ioutil.ReadAll(resp.Body)
					//if err != nil {
					//	msg.Text = err.Error()
					//	SendMessage(msg)
					//}
					//bytes := tgbotapi.FileBytes{Name: "image.jpg", Bytes: b}
					photo := tgbotapi.NewInputMediaPhoto(tgbotapi.FileURL(item))
					mediaGroup.Media = append(mediaGroup.Media, photo)
					if len(mediaGroup.Media) >= 9 {
						mediaGroups = append(mediaGroups, mediaGroup)
						mediaGroup.Media = []interface{}{}
					}
					if len(mediaGroup.Media) < 9 && index == len(chanMsg.ImgList)-1 {
						mediaGroups = append(mediaGroups, mediaGroup)
					}
				}
				for _, item := range mediaGroups {
					_, err := instance.SendMediaGroup(item)
					if err != nil {
						msg.Text = err.Error()
						glog.Errorf("SendMediaGroup error %s", err.Error())
						SendMessage(msg)
					}
					time.Sleep(500 * time.Millisecond)
				}
			}

		}
		glog.Info("已退出")
	}(recMsgChan)
	// 初始化订阅
	subscribe := redisTool.Subscribe{SubscribeChannel: "channel"}
	subscribe.New(recMsgChan)
}
