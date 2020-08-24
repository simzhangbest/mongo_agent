package web

import (
	"github.com/gin-gonic/gin"
)

/*
	定义路由地址和对应路由地址的响应函数
 */

var appinstance *webapp		//webapp单例对象
type webapp struct {
	server *gin.Engine
}



func init(){
	appinstance = new(webapp)
	appinstance.server = gin.Default()		//获取web服务器对象
	//server.LoadHTMLGlob("Server/templates")		//加载静态页面
}

func Newwebapp() *webapp{
	return appinstance
}

func (app *webapp)Prepare()  {
	app.server.GET("/do", IndexRouter)		//创造一个GET请求的路由地址，并指定处理函数
	app.server.POST("/userMoods", UserInfoDeal)
	app.server.POST("/userEditorOpr", UserEditorOpr)
	app.server.POST("/userBlur", UserBlur)
	app.server.POST("/userFocus", UserFocus)
	app.server.POST("/userClose", UserClose)
	app.server.GET("/ping", UserPong)

	// for pa audio using
	app.server.GET("/paHeartBeat", PaHB)
	app.server.GET("/paPlay", PaPlay)
	app.server.GET("/paPause", PaPause)
	app.server.GET("/paStop", PaStop)

	app.server.GET("/ws", GinWebsocketHandler(WsConnHandle))
	//app.server.POST("/do", IndexRouter)
	app.server.Run(":8088")
}





