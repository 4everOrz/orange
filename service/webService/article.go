package webService

import (
	"errors"
	"github.com/labstack/echo"
	"net/http"
	"orange/controller/article"
	"orange/controller/user"
	"orange/runtime"
)

type articleHandler struct {
	runtime runtime.Runtime
}

func (this *articleHandler)Add(ctx echo.Context)(err error)  {
	cmd:=&article.AddCmd{}
	if err:=ctx.Bind(cmd);err!=nil{
		err=errors.New("参数异常")
		goto End
	}

	cmd.Creator=(ctx.Get("userinfo")).(*user.User).ID
	if err=this.runtime.Article().Add(cmd);err!=nil{
		goto End
	}
	return ctx.JSON(http.StatusOK, response{Message: "操作成功"})
End:
	return ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
}