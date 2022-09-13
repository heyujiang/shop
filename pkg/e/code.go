package e

import "fmt"

//go:generate stringer -type Code
type Code int

const (
	Success Code = 200 //成功
	Fail    Code = 400

	BodyQuestErr     Code = 401 //Body中的请求参数错误
	ParamQuestErr    Code = 402 //Param中的请求参数错误
	ResourceNotFound Code = 404 //资源未发现
)

func (i Code) ToDesc() string {
	switch i {
	case Success:
		return "请求成功"
	case Fail:
		return "请求失败"
	case BodyQuestErr:
		return "Body中的请求参数错误"
	case ParamQuestErr:
		return "Param中的请求参数错误"
	case ResourceNotFound:
		return "资源未发现"
	}
	return fmt.Sprintf("%d", i)
}
