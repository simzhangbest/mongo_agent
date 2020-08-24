package main

import "example.com/m/Server/web"

//import "Server/web"

/*
	程序入口
*/

func main() {
	webapp := web.Newwebapp()
	webapp.Prepare()
}



