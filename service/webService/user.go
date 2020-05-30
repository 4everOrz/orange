package webService

import (
	"errors"
	"net/http"
	"orange/common/token"
	"orange/controller/user"
	"orange/runtime"
	"security"
	"strings"

	"github.com/alecthomas/log4go"
	"github.com/labstack/echo"
)

type userHandler struct {
	runtime runtime.Runtime
}

func (this *userHandler) Add(ctx echo.Context) (err error) {
	cmd := &user.AddCmd{}

	if err = ctx.Bind(cmd); err != nil {
		log4go.Error(err.Error())
		goto End
	}

	if err = this.runtime.User().Add(cmd); err != nil {
		log4go.Error(err.Error())
		goto End
	}
	return ctx.JSON(http.StatusOK, response{Message: "操作成功"})
End:
	return ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
}
func (this *userHandler)Get(ctx echo.Context)(err error){
	res:=&user.User{}

	id:=ctx.Param("id")
	if strings.TrimSpace(id)==""{
		goto End
	}

	res,err=this.runtime.User().Get(id,"")
	if err!=nil{
		log4go.Error(err.Error())
		goto End
	}
	return ctx.JSON(http.StatusOK, response{Message: "操作成功",Data: res})
End:
	return ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
}
func (this *userHandler)Me(ctx echo.Context)(err error){
	return ctx.JSON(http.StatusOK, response{Message: "操作成功",Data: ctx.Get("userinfo")})
}
func (this *userHandler)Delete(ctx echo.Context)(err error){
	 var ids struct{
		IDs []string `json:"ids"`
	}
	if err=ctx.Bind(&ids);err!=nil{
		err=errors.New("参数解析异常")
		goto End
	}

	if err=this.runtime.User().Delete(ids.IDs);err!=nil{
		log4go.Error(err.Error())
		goto End
	}
	return ctx.JSON(http.StatusOK, response{Message: "操作成功"})
End:
	return ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
}
func (this *userHandler)Query(ctx echo.Context)(err error){
	var resp user.UserList
	cmd:= &user.QueryCmd{}
	if err=ctx.Bind(cmd);err!=nil{
		err=errors.New("参数解析异常")
		goto End
	}
    resp,err= this.runtime.User().Query(cmd)
    if err!=nil{
    	goto End
	}
	return ctx.JSON(http.StatusOK, response{Message: "操作成功",Data: resp})
End:
	return ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
}

type authResp struct{
	Token string `json:"token"`
	UserID string  `json:"id"`
	Role   int     `json:"role"`
}

func (this *userHandler)Auth(ctx echo.Context)(err error){
    userEntity:=&user.User{}
    resp:=&authResp{}
    var auth struct{
    	LoginName  string  `json:"loginName"`
    	Password   string  `json:"password"`
	}
	//绑定参数
	if err=ctx.Bind(&auth);err!=nil{
		err=errors.New("参数解析异常")
		goto End
	}
	//获取用户信息
	userEntity,err=this.runtime.User().Get("",auth.LoginName)
	if err!=nil{
		log4go.Error(err.Error())
		err=errors.New("获取用户信息失败")
		goto End
	}
	if userEntity!=nil{
		//判断用户使能
		if userEntity.Enable==1{
			err=errors.New("用户已禁用")
			goto End
		}
		//校验密码
		if userEntity.Password==security.Md5(auth.Password){
			resp.Token=token.GenToken(userEntity.ID)
			resp.UserID=userEntity.ID
			resp.Role=userEntity.Role
		}else {
			err=errors.New("用户名或密码错误")
			goto End
		}
	}else {
		err=errors.New("用户名或密码错误")
		goto End
	}
	return ctx.JSON(http.StatusOK, response{Message: "用户验证通过",Data:resp})
End:
	return ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
}
func (this *userHandler)Update(ctx echo.Context)(err error)  {
	var userEntity *user.User
	 cmd :=&user.UpdateCmd{}

	if err=ctx.Bind(cmd);err!=nil{
		err=errors.New("参数解析异常")
		goto End
	}
	userEntity=ctx.Get("userinfo").(*user.User)
	cmd.ID=userEntity.ID
	err=this.runtime.User().Update(cmd)
	if err!=nil{
		log4go.Error(err.Error())
		goto End
	}
	return ctx.JSON(http.StatusOK, response{Message: "操作成功"})
End:
	return ctx.JSON(http.StatusBadRequest, response{Message: err.Error()})
}