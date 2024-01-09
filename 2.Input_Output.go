package main

import "fmt"

func main() {

	//输入
	fmt.Println("-----接下来是测试输入-----")
	fmt.Printf("请输入你的名字：")
	var nameInput string
	fmt.Scan(&nameInput)
	fmt.Printf("\n请输入你的年龄：")
	var ageInput int
	fmt.Scan(&ageInput)
	fmt.Printf("你的名字是%s, 你今年%d岁\n", nameInput, ageInput)

	//输出
	fmt.Println("-----接下来是测试输出-----")
	fmt.Println("HI")
	fmt.Println("hello world", "你好")
	//格式化输出
	var name = "Xie"
	fmt.Printf("哇, %s，是你啊\n", name)
	fmt.Printf("%d\n", 3)
	fmt.Printf("%f\n", 3.1415926)
	fmt.Printf("%.2f\n", 3.1415926)
	fmt.Printf("%T %T\n", 123, "HI") //%T：打印类型
	fmt.Printf("%v\n", 123)          //%v 不管什么，打印变量

}
