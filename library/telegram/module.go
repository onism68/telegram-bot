package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strings"
	"sync"
)

var (
	Modules      = make(map[string]ModuleInfo)
	ModulesMu    sync.RWMutex
	GlobalModule = "global"
	CronModule   = "cron"
	ActiveCron   = "Active"
)

type Module interface {
	ModuleInfo() ModuleInfo

	//module 生命周期

	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
	Init()

	// PostInit 第二次初始化
	// 调用该函数时，所有 Module 都已完成第一段初始化过程
	// 方便进行跨Module调用
	PostInit()

	// Serve 向Bot注册服务函数
	// 结束后调用 Start
	Serve(bot *Bot)

	// Start 启用Module
	Start(bot *Bot, update tgbotapi.Update)

	// Stop 应用结束时对所有 Module 进行通知
	// 在此进行资源回收
	Stop(bot *Bot, wg *sync.WaitGroup)
}

// id
type ModuleId string

// Namespace - 获取一个 Module 的 Namespace
func (id ModuleId) Namespace() string {
	lastDot := strings.LastIndex(string(id), ".")
	if lastDot < 0 {
		return ""
	}
	return string(id)[:lastDot]
}

// Name - 获取一个 Module 的 Name
func (id ModuleId) Name() string {
	if id == "" {
		return ""
	}
	parts := strings.Split(string(id), ".")
	return parts[len(parts)-1]
}

// 模块信息
type ModuleInfo struct {
	Id       ModuleId
	Instance Module
}

// 全局module注册
func RegisterModule(instance Module) {
	module := instance.ModuleInfo()

	if module.Id == "" {
		panic("module Id is empty")
	}

	if module.Instance == nil {
		panic("module instance is nil")
	}

	ModulesMu.Lock()
	defer ModulesMu.Unlock()
	// 判断module是否已经注册
	if _, ok := Modules[string(module.Id)]; ok {
		panic(fmt.Sprintf("module already registered: %s", module.Id))
	}
	Modules[string(module.Id)] = module
}

// GetModule - 获取一个已注册的 Module 的 ModuleInfo
func GetModule(name string) (ModuleInfo, error) {
	ModulesMu.Lock()
	defer ModulesMu.Unlock()
	m, ok := Modules[name]
	if !ok {
		return ModuleInfo{}, fmt.Errorf("module not registered: %s", name)
	}
	return m, nil
}
