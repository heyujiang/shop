package sys

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {

}

func GetUserInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println(id)
}
