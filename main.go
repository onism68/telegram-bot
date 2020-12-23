package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	_ "telegram-bot/boot"
	"telegram-bot/library/telegram"
	_ "telegram-bot/library/telegram/modules/help"
	_ "telegram-bot/library/telegram/modules/logging"
	_ "telegram-bot/library/telegram/modules/myIp"
	_ "telegram-bot/library/telegram/modules/repeater"
	_ "telegram-bot/router"
)

func main() {
	bot, err := telegram.BotInit()
	if err != nil {
		panic(err)
	}
	glog.Infof("bot [%s] 已登录", bot.Self.UserName)
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, os.Interrupt, os.Kill)
	//<-ch
	//bot.Stop()
	//<-telegram.Running
	g.Server().Run()

}
