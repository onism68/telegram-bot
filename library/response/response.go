package response

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var (
	Ok            string = "操作成功"
	InputEmpty    string = "请完整填写相关项目"
	DbError       string = "数据库错误，请联系管理员"
	ServiceError  string = "服务出错，请联系管理员"
	PasswordError string = "用户名或者密码错误"
	RequestError  string = "请求错误"
)

// 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// err:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;

//10月14日， 将err错误码更改为code状态码

func Json(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  msg,
		"data": responseData,
	})
	r.Exit()
}
