package main

import "fmt"

func main() {

	var age int
	fmt.Println("输入一个年龄")
	fmt.Scan(&age)

	//switch，自带break的，用fallthrough往下走
	switch {
	case age < 18:
		fmt.Println("小孩")
	case age < 35:
		fmt.Println("好好干")
	case age < 55:
		fmt.Println("多休息")
		fallthrough
	default:
		fmt.Println("没什么好说的")

	}

	//switch 另一种写法
	switch age {
	case 10:
		fmt.Println("生活元年")
	case 20:
		fmt.Println("生活元年")
	case 30:
		fmt.Println("生活元年")
	}

}
