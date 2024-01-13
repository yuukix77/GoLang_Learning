package main

import "fmt"

//空接口相当于所有的实例都被自动继承

type EmptyInterface interface{}

func Print1(val EmptyInterface) {
	fmt.Println(val)
	fmt.Println("使用空接口")
}

//另外的写法-1

func Print2(val interface{}) {
	fmt.Println(val)
	fmt.Println("使用空接口-2")
}

// 另外的写法-2
func Print3(val any) {
	fmt.Println(val)
	fmt.Println("使用空接口-3")
}

func main() {
	testVal := 1
	testVal2 := "HI"
	Print1(testVal)
	Print1(testVal2)

	Print2(testVal)
	Print2(testVal2)

	Print3(testVal)
	Print3(testVal2)

}
