package controler

import (
	"example.com/m/Server/db"
	"example.com/m/Server/utils"
	"fmt"

	//"Server/db"
	"time"
	//"Server/utils"
	//"../utils"
)


var dbop *db.Operdb //数据库操作对象
var uid_generator *utils.Uid

func init() {
	dbop = db.Newoperdb()
	uid_generator = utils.Newuid()
}

type Handle struct {
	count int		//对未插入数据库的请求进行计数
}

//返回初始化对象，并对对象初始化
func NewHandle()  *Handle{
	handle := new(Handle)
	handle.count = 0
	go handle.gorount()
	return handle
}

//加上标志字段，将参数传递给数据库操作类进行插入
func (handle *Handle)Insert_ser(arg map[string][]string) {
	arg1 := make(map[string]interface{})		//用于将map[string][]string 转换成 map[string]string
	for k, v := range arg{
		arg1[k] = v[0]
	}
	id := uid_generator.Getid()			//获取唯一流水线标志
	arg1["insertid"] = id
	dbop.BulkInset(arg1)
	handle.count++
	fmt.Println(handle.count)
}

// sim add 学生所用类型的数据插入
// 格式如下：

func (handle *Handle)Insert_user_data  (arg interface{}) {
	fmt.Println("sim deal mongodb")
	id := uid_generator.Getid()			//获取唯一流水线标志
	arg1 := make(map[string]interface{})
	arg1["insertid"] = id
	dbop.BulkInset(arg1)
	handle.count++
}

//协程，定期或定量向数据库批量插入
//此处代码的意思 时间间隔由后端处理
// 100000000   100 ms
// 100         count 计数 100次
func (handle *Handle)gorount()  {
	for{
		starttime := time.Now().UnixNano()
		for{
			endtime := time.Now().UnixNano()
			interval := endtime - starttime		//得到间隔时间
			if interval >= 100000000 || handle.count >= 100{	//如果间隔超过100毫秒或者数量超过100就插入数据库
				break
			}
		}
			dbop.BulkRun()
			handle.count = 0
	}
}