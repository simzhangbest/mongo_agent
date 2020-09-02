package web

//import "Server/controler"
import (
	"example.com/m/Server/controler"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/websocket"
	"net/http"
)
/*
	定义路由处理函数
*/
var handle *controler.Handle   //逻辑处理对象
var flag string
//pingInfo := make(chan string)

func init()  {
	handle = controler.NewHandle()
}

func IndexRouter(c *gin.Context) {
	if c.Request.Form == nil {		//获取所有请求参数名和值
		c.Request.ParseMultipartForm(32 << 20)
	}

	handle.Insert_ser(c.Request.Form)

	//c.HTML(http.StatusOK, "index.html", nil)		//页面跳转
}


// 处理心情
func UserInfoDeal(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")

	// 处理返回值
	c.JSON(http.StatusOK, gin.H{
		"method": c.Request.Method,
		"moods": c.PostForm("moods"),
		"username": c.PostForm("username"),
		"time": c.PostForm("time"),
	})

	c.Request.ParseForm()
	handle.Insert_ser(c.Request.PostForm)
	// 组成一个map
	//moodsData := make(map[string] interface{})
	//wg := sync.WaitGroup{}
	//moodsData = {
	//	"method": c.Request.Method,
	//		"moods": c.PostForm("moods"),
	//		"username": c.PostForm("username"),
	//		"time": c.PostForm("time"),
	//}
	//handle.Insert_user_data()

}

// 处理编辑器
func UserEditorOpr(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	//fmt.Println("editor")
	//for k, v := range c.Request.PostForm {
	//	fmt.Printf("k: %v  ", k)
	//	fmt.Printf("v: %v\n", v)
	//}

	c.Request.ParseForm()
	handle.Insert_ser(c.Request.PostForm)
}

// 处理blur
func UserBlur(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	//fmt.Println("editor")
	c.Request.ParseForm()
	handle.Insert_ser(c.Request.PostForm)
}


// 处理focus
func UserFocus(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	//fmt.Println("editor")
	c.Request.ParseForm()
	handle.Insert_ser(c.Request.PostForm)
}

// 处理关闭
func UserClose(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	//fmt.Println("editor")
	c.Request.ParseForm()
	handle.Insert_ser(c.Request.PostForm)
}


func UserPong(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"message": "pong",
	})

	stuId :=c.Request.Header.Get("STU_ID")

	fmt.Println("STU_ID IS: ", stuId)

	fmt.Println("I am Pong")
	msg := c.Request.RemoteAddr
	pinInfo := msg + "+" + " sim ping"
	if tmpWsCon != nil {
		simWSDataWriter(pinInfo, tmpWsCon)
	}
	//simWsDataWrite(pingInfo)
}

func simWSDataWriter(info string, conn *websocket.Conn) {
	wsconnectedwith := conn.RemoteAddr().String()
	fmt.Println("ws client is: \n" +  wsconnectedwith)
	data := []byte(info)
	if _, err := conn.Write(data); err != nil {
		fmt.Println(err)
		return
	}
}

func PaStop(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"message": "stop",
	})


	stuId :=c.Request.Header.Get("STU_ID")

	fmt.Println("STU_ID IS: ", stuId)
	msg := c.Request.RemoteAddr
	pinInfo := msg + "+" + "stop" + "+" + stuId
	if tmpWsCon != nil {
		simWSDataWriter(pinInfo, tmpWsCon)
	}
}

func PaPause(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"message": "pause",
	})

	msg := c.Request.RemoteAddr
	pinInfo := msg + "+" + "pause"
	if tmpWsCon != nil {
		simWSDataWriter(pinInfo, tmpWsCon)
	}
}

func PaPlay(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"message": "play",
	})
	stuId :=c.Request.Header.Get("STU_ID")
	fmt.Println("STU_ID IS: ", stuId)

	msg := c.Request.RemoteAddr
	pinInfo := msg + "+" + "play" + "+" + stuId
	if tmpWsCon != nil {
		simWSDataWriter(pinInfo, tmpWsCon)
	}

}


// ws 链接状态即可判断，不需要心跳
func PaHB(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, gin.H{
		"message": "heart beat",
	})


	//msg := c.Request.RemoteAddr
	//pinInfo := msg + "+" + "HB"
	//if tmpWsCon != nil {
	//	simWSDataWriter(pinInfo, tmpWsCon)
	//}

}

func GinWebsocketHandler(wsConnHandle websocket.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("new ws request: %v\n", c.Request.RemoteAddr)
		if c.IsWebsocket() {
			wsConnHandle.ServeHTTP(c.Writer, c.Request)
		} else {
			_, _ = c.Writer.WriteString("===not websocket request===")
		}

	}
}


// 此处只做接收数据显示，和引出ws 全局变量
// 此处for 循环 需要优化一下，不然会导致cpu 线程跑满
var tmpWsCon *websocket.Conn
func WsConnHandle(conn *websocket.Conn) {
	tmpWsCon = conn
	fmt.Println(tmpWsCon)
	//for {
	//	var msg string
	//	if err := websocket.Message.Receive(conn, &msg); err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//
	//	fmt.Printf("recv: %v time: %v from remote addr: %v \n", msg, time.Now(), conn.RemoteAddr().String())

		//data := []byte(time.Now().Format(time.RFC3339) + " simzhangtest")
		//data := []byte(flag)

		// heartbeat
		//for {
		//	time.Sleep(3 * time.Second)
		//	fmt.Println(flag)
		//	if _, err := conn.Write(data); err != nil {
		//		fmt.Println(err)
		//		return
		//	}
		//}
		//}
}