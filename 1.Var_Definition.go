package main

import (
	version "GoSyntaxStudy/Packet_Reference_Test"
	"fmt"
)

/*
注：函数内声明的变量全是局部变量，而且智能用var定义,不可以采用简单声明方式，不受使用与否的规律限制
*/
var globalVar = "我是全局变量globalVar"

// 多变量定义-1
var globalVar1, globalVar2 = 1, 2

// 多变量定义-2
var (
	globalVar3 = 4
	globalVar4 = 5
)

// 常量
const constant = "我是常量"

func hello() {
	fmt.Println("This is function 'hello()'")
}

func main() {
	//变量定义方式-1
	var name string
	name = "Xie"
	fmt.Println(name)

	//变量定义方式-2
	var name2 string = "Xie2"
	fmt.Println(name2)

	//变量定义-3 省略类型
	var name3 = "Xie3"
	fmt.Println(name3)

	//变量定义-4 省略类型+省略var
	name4 := "Xie4"
	fmt.Println(name4)

	//测试一些用法
	hello()
	fmt.Println(globalVar)
	fmt.Println(constant)
	fmt.Println("跨包访问")
	fmt.Println(version.Version)

}
