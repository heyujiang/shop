package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/heyujiang/shop/app/web/server/routes/util"
	"github.com/heyujiang/shop/model/req"
	"github.com/heyujiang/shop/pkg/e"
)

func Login(ctx *gin.Context) {
	loginReq := req.Login{}
	if err := ctx.BindJSON(&loginReq); err != nil {
		util.SendResponse(ctx, e.BodyQuestErr, nil)
	}

}

func GetUserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
}
