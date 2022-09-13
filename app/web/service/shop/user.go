package shop

import (
	"context"
	commonV1 "github.com/heyujiang/hapis/protogen-go/common/v1"
	userV1 "github.com/heyujiang/hapis/protogen-go/user/v1"
	"github.com/heyujiang/shop/app/web/service/user_client"
	"github.com/heyujiang/shop/model/req"
	"github.com/heyujiang/shop/model/view"
	"github.com/heyujiang/shop/pkg/e"
)

func Login(req req.Login) (e.Code, view.LoginResp) {
	result, err := user_client.GetUserClient().Login(context.Background(), &userV1.UserLoginReq{
		UserName: req.UserName,
		PassWord: req.PassWd,
	})
	if err != nil {

	}

	if result.Code != commonV1.StatusCode_Success {

	}

	return e.Success, view.LoginResp{UserName: req.UserName}
}
