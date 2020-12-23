package myIp

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *MyIp

func init() {
	//instance := &MyIp{}
	//telegram.RegisterModule(instance)
}

type MyIp struct{}

func (myIp *MyIp) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       "test.myIp",
		Instance: instance,
	}
}

func (myIp *MyIp) Init() {
	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
	//panic("implement me")
}

func (myIp *MyIp) PostInit() {
	// 第二次初始化
	// 再次过程中可以进行跨Module的动作
	// 如通用数据库等等
	//panic("implement me")
}

func (myIp *MyIp) Serve(bot *telegram.Bot) {
	// 注册服务函数部分
	//panic("implement me")
}

func (myIp *MyIp) Start(bot *telegram.Bot, update tgbotapi.Update) {
	// 此函数会新开协程进行调用
	// ```go
	// 		go exampleModule.Start()
	// ```
	//panic("implement me")
}

func (myIp *MyIp) Stop(bot *telegram.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	defer wg.Done()
	// 结束部分
	// 一般调用此函数时，程序接收到 os.Interrupt 信号
	// 即将退出
	// 在此处应该释放相应的资源或者对状态进行保存
	//panic("implement me")
}
