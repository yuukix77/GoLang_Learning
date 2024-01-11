package main

import (
	"fmt"
)

func main() {
	var age int

	fmt.Printf("现在，age等于%v\n", age)
	for age = 10; age < 100; age++ {
		fmt.Println("age++....")
	}
	fmt.Printf("现在，age等于%v\n", age)

	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	////死循环
	//for true {
	//	fmt.Println(time.Now())
	//	time.Sleep(1 * time.Second)
	//}

	sum = 0
	for i := 1; i <= 200; i++ {
		if i == 50 {
			continue
		}
		sum += i
		if i == 100 {
			break
		}
	}
	println(sum)

	var listName = []string{"xie", "yang", "ruan"}

	for i, name := range listName {
		fmt.Println(i, name)
	}

	var locationMap = map[int]string{
		1: "wuhu",
		2: "guangzhou",
		3: "changsha",
	}
	for key, value := range locationMap {
		fmt.Println(key, value)
	}
}
