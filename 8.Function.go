package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func sayHello2(name string) {
	fmt.Printf("Hi,%v\n", name)
}

func sayHello3(num int, name string) {
	fmt.Printf("第%v位新员工, %v, 您好，欢迎您\n", num, name)
}

// 多参数
func sum(numList ...int) {
	var sumOfnumList int
	for _, i := range numList {
		sumOfnumList += i
	}
	fmt.Println(sumOfnumList)
}

func sum2(numList ...int) int {
	var sumOfnumList int
	for _, i := range numList {
		sumOfnumList += i
	}
	return sumOfnumList
}

// 多返回值
func sum3(numList ...int) (int, int) {
	var sumOfnumList int
	for _, i := range numList {
		sumOfnumList += i
	}
	return len(numList), sumOfnumList
}

// 返回值提前定义
func sum4(numList ...int) (lenOfnumList int, sumOfnumList int) {
	lenOfnumList = len(numList)
	for _, i := range numList {
		sumOfnumList += i
	}
	return
}

func login() {
	println("登录")
}

func register() {
	println("注册")
}

func personCenter() {
	println("个人中心")
}

func main() {
	sayHello()
	sayHello2("XiaoXie")
	sayHello3(3, "Wang Di")
	sum(1, 2, 3)
	var testSum2 = sum2(1, 2)
	println(testSum2)

	var lenOfList, testSum3 = sum3(2, 4, 6)
	fmt.Printf("总共有%d个数，加起来是%d\n", lenOfList, testSum3)

	var lenOfList4, testSum4 = sum4(1, 3)
	fmt.Printf("总共有%d个数，加起来是%d\n", lenOfList4, testSum4)

	//匿名函数
	var getName = func() string {
		return "XiaoXie"
	}
	fmt.Println(getName())

}
