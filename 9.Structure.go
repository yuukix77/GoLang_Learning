package main

import "fmt"

type Class struct {
	className string
}

type Student struct {
	name string
	id   int
	Class
}

func Study(stu Student) {
	fmt.Printf("%v正在学习\n", stu)
}

func main() {
	var class7 = Class{className: "7班"}
	var stu1 = Student{name: "Xie", id: 1001, Class: class7}
	Study(stu1)

}
