package user

import (
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) interface{}{
	userName,isExist:= ctx.Get("userName")
	if isExist{
		return userName
	}

	return nil
}
