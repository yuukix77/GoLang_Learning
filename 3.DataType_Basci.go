package main

import "fmt"

func main() {

	//整型
	var unsignedInt1 uint = 255
	var unsignedInt2 uint8 = 255
	var Int2 int //未定义就等于0
	fmt.Println(unsignedInt1, unsignedInt2, Int2)

	//浮点型
	var f1 float32 = 3.14
	var f2 float64 = 3.1415926 //默认是64
	fmt.Println(f1, f2)

	//字符型和字符串型
	var a byte = 'a' //ascii
	var a1 int = 97
	fmt.Printf("%c %d\n", a, a)
	fmt.Printf("%c %d\n", a1, a1)

	var z rune = '中' //非英语语言，unicode, rune：符文
	fmt.Printf("%c %d\n", z, z)

	var s string = "Hello"
	fmt.Printf("s=%v, 类型是：%T, 字符串长度为：%d\n", s, s, len(s))

	//布尔类型
	//Go不允许布尔类型转其他类型，也就是0!=False
	var b1 bool = true
	var b2 bool = false
	fmt.Println(b1, b2)

	//零值问题,不赋值的默认值
	var O int
	var O1 float64
	var O2 bool
	var O3 string
	fmt.Printf("%#v\n", O)
	fmt.Printf("%#v\n", O1)
	fmt.Printf("%#v\n", O2)
	fmt.Printf("%#v\n", O3)
}
