package webService

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	_ "net/http/pprof"
	"orange/common/token"
	_ "orange/controller"
	"orange/runtime"
	"orange/service"
	"strings"
)

type WebService interface {
        service.Service
}

type webService struct {
 listenPort string
 engine  *echo.Echo
 runtime  runtime.Runtime
}
type response struct {
	Message string `json:"message,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}
func New(port string,runtime runtime.Runtime)WebService{
	return &webService{
		listenPort: port,
		engine: echo.New(),
		runtime: runtime,
	}
}
func (this *webService)Name()string{
	return "webService"
}

func (this *webService)Init(){

}
func (this *webService)Start()error{
	//pprof工具
	go func(){
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()
	//web设置
	this.setMiddlewares()
	this.setRoutes()
	this.engine.Start(":"+this.listenPort)
	//会阻塞
	return nil
}
func (this *webService)Stop()error  {
	return this.engine.Close()
}
//设置中间件
func (this *webService)setMiddlewares(){
	this.engine.Use(middleware.CORS())
	this.engine.Use(middleware.Gzip())
	this.engine.Use(middleware.Recover())
	this.engine.Use(middleware.BodyLimit("16M"))
	this.engine.Use(middleware.Static("view/static"))
}
//设置路由
func (this *webService)setRoutes(){
	//静态文件
	this.engine.Static("/static", "view/static")
	this.engine.File("/", "view/index.html")
	//api路由
	version := this.engine.Group("/v1")
	{
		userHandler:=userHandler{this.runtime}
		user := version.Group("/users")
		user.POST("", userHandler.Add)
		user.GET("/:id",userHandler.Get,needAdmin(this.runtime))
		user.DELETE("",userHandler.Delete,needAdmin(this.runtime))
		user.GET("",userHandler.Query,needAdmin(this.runtime))
		user.PUT("/auth",userHandler.Auth)
		user.GET("/me",userHandler.Me,userInfo(this.runtime))
		user.PUT("/:id",userHandler.Update,userInfo(this.runtime))
		user.GET("/:id/all",userHandler.QueryAll)
	}
   {
	   articleHandler:=articleHandler{this.runtime}
	   article:=version.Group("/articles")
	   article.POST("",articleHandler.Add,userInfo(this.runtime))
    }

	//{
	//	userCtrl := controllers.UserController{}
	//	user := version.Group("/users")
	//	user.POST("/", userCtrl.AddOne)                 //新增
	//	user.POST("/login", userCtrl.Login)             //登录
	//}
}

func userInfo(runtime runtime.Runtime) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			tokenstr := ectx.Request().Header.Get("authorization")
			tokenstr = strings.TrimSpace(tokenstr)
			if len(tokenstr) == 0 {
				return ectx.JSON(http.StatusBadRequest,response{Message: "用户未登录"})
			}
           userID,ok:=token.CheckToken(tokenstr)
           if !ok{
          	 return ectx.JSON(http.StatusBadRequest,response{Message: "用户token检验失败"})
		   }
		  userInfo,err:= runtime.User().Get(userID,"")
		  if err!=nil{
			 return ectx.JSON(http.StatusBadRequest,response{Message: "获取用户失败"})
		 }
			ectx.Set("userinfo",userInfo)
			return next(ectx)
		}
	}
}

func needAdmin(runtime runtime.Runtime) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			tokenstr := ectx.Request().Header.Get("authorization")
			tokenstr = strings.TrimSpace(tokenstr)
			if len(tokenstr) == 0 {
				return ectx.JSON(http.StatusBadRequest,response{Message: "用户未登录"})
			}
			userID,ok:=token.CheckToken(tokenstr)
			if !ok{
				return ectx.JSON(http.StatusBadRequest,response{Message: "用户token检验失败"})
			}
			userInfo,err:= runtime.User().Get(userID,"")
			if err!=nil||userInfo==nil{
				return ectx.JSON(http.StatusBadRequest,response{Message: "获取用户失败"})
			}
		    if userInfo.Role!=1{
		 	return ectx.JSON(http.StatusBadRequest,response{Message: "用户未授权"})
		   }
			return next(ectx)
		}
	}
}
