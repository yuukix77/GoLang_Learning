package main

import "fmt"

//defer函数，离return近的先执行

func main() {

	defer fmt.Println("1")
	defer fmt.Println("2")
	return
}
