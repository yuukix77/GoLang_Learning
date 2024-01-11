package main

import "fmt"

func Copy(name string) {
	fmt.Printf("%p\n", &name)
}

func CopyWithPointer(name *string) {
	fmt.Printf("%p\n", name)
	*name = "牛的"
}

func main() {
	var name = "Xie"
	fmt.Printf("%p\n", &name)
	Copy(name)
	fmt.Println(name)
	CopyWithPointer(&name)
	fmt.Println(name)
}
