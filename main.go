package main

import userV1 "github.com/heyujiang/hapis/protogen-go/user/v1"

func main() {
	_ = userV1.NewUserClient(nil)
}
