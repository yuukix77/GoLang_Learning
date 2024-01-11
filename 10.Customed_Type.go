package main

import "fmt"

type Code int

const (
	Success_Code  Code = 0
	Network_Error Code = 1001
	Service_Error Code = 1002
)

func (c Code) ServiceSystem() (msg string) {
	switch c {
	case Success_Code:
		msg = "连接成功"
		return
	case Network_Error:
		msg = "网络错误"
		return
	case Service_Error:
		msg = "服务器错误"
		return
	default:
		return
	}
}

func main() {
	var code = Success_Code
	var code2 = Service_Error
	fmt.Println(code.ServiceSystem())
	fmt.Println(code2.ServiceSystem())
}
