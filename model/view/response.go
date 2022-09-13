package view

import "github.com/heyujiang/shop/pkg/e"

type Response struct {
	Success bool
	Code    e.Code
	Msg     string
	Data    interface{}
}
