package main

import "fmt"

func main() {
	var age int
	fmt.Println("输入一个年龄")
	fmt.Scan(&age)

	if age <= 18 {
		fmt.Println("你是未成年")

	}
	if age > 18 {
		fmt.Println("你是成年人")
	} else if age < 5 {
		fmt.Println("一个小孩")
	} else {
		fmt.Println("没关系")
	}

	if age >= 35 && age <= 60 {
		fmt.Println("唉，小心被裁")
	}

}
