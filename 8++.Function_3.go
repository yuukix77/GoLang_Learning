package main

import "fmt"

//特殊函数init

func init() {
	fmt.Println("init-1")
}
func init() {
	fmt.Println("init-2")
}
func init() {
	fmt.Println("init-3")
}

func main() {
	fmt.Println("main")
}
