package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Resolve() map[string]string{
	value := make(map[string]string)
	// 注意此处，win 和 linux 的路径使用不一样
	str, _ := os.Getwd()
	fmt.Println(str)
	// windows 读取有问题
	b ,err :=ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	json.Unmarshal(b,&value)


	return value
}

