package redisTool

import (
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"telegram-bot/library/telegram/types"
)

type Subscribe struct {
	SubscribeChannel string
}

func (s *Subscribe) New(msgChan chan types.TgMsg) {

	redis, err := gredis.NewFromStr(g.Cfg().GetString("redis.default"))
	conn := redis.Conn()
	//conn := g.Redis().Conn()
	//defer conn.Close()
	_, err = conn.Do("SUBSCRIBE", s.SubscribeChannel)
	if err != nil {
		glog.Errorf("redisTool do {SUBSCRIBE %s} error of %s", s.SubscribeChannel, err)
		return
	}
	for {
		resp, err := conn.ReceiveVar()
		if err != nil {
			glog.Errorf("conn receive error of %s", err.Error())
			return
		}
		// 将收到的消息加到msgChan中
		var tgMsg *types.TgMsg
		err = gjson.DecodeTo(resp.Array()[2], &tgMsg)
		if err != nil {
			glog.Errorf("conn receive error of %s", err.Error())
			return
		}
		glog.Debug(tgMsg)
		//if err = resp.Struct(&tgMsg); err != nil {
		//	glog.Errorf("conv struct error of %s", err.Error())
		//	return
		//}
		msgChan <- *tgMsg
	}
}
