package user_client

import (
	"context"
	common "github.com/heyujiang/hapis/protogen-go/common/v1"
	user "github.com/heyujiang/hapis/protogen-go/user/v1"
	"github.com/heyujiang/shop/config"
	"github.com/heyujiang/shop/model"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"sync"
)

var initClientUserOnce sync.Once
var userClient user.UserClient

func getUserClient() user.UserClient {
	initClientUserOnce.Do(func() {
		conn, _ := grpc.Dial(config.GetConfig().User.Url, grpc.WithInsecure())
		userClient = user.NewUserClient(conn)
	})
	return userClient
}

func Login() error {
	resp, err := userClient.Login(context.Background(), &user.UserLoginReq{
		UserName: "",
		PassWord: "",
	})
	if err != nil {
		return errors.Wrap(err, "Login fail")
	}
	if resp.GetCode() != common.StatusCode_Success {
		return errors.Errorf("Get user info fail : %s", resp.GetMsg())
	}
	return nil
}

func GetUserInfo() (model.User, error) {
	userInfo := model.User{}
	resp, err := userClient.GetUserInfo(context.Background(), &user.GetUserInfoReq{Id: 1})
	if err != nil {
		return userInfo, errors.Wrap(err, "Get user info fail")
	}
	if resp.GetResult().GetCode() != common.StatusCode_Success {
		return userInfo, errors.Errorf("Get user info fail : %s", resp.GetResult().GetMsg())
	}

	userInfo.Id = resp.GetUserInfo().GetId()
	userInfo.Name = resp.GetUserInfo().GetName()
	return userInfo, nil
}
