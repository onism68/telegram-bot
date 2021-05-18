package myIp

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gogf/gf/os/glog"
	"io/ioutil"
	"net/http"
	"sync"
	"telegram-bot/library/telegram"
)

var instance *MyIp
var id = "cmd.myIp"

func init() {
	glog.Debugf("开始注册: %s", id)
	telegram.RegisterModule(instance)
}

type MyIp struct{}

func (myIp *MyIp) ModuleInfo() telegram.ModuleInfo {
	return telegram.ModuleInfo{
		Id:       telegram.ModuleId(id),
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
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	resp, err := http.Get("http://api.ip.sb/ip")
	if err != nil {
		msg.Text = err.Error()
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	msg.Text = string(b)
	telegram.SendMessage(msg)
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
