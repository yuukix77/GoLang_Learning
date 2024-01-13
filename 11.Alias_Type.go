package main

import "fmt"

type NetCode int        //类型定制
type AliasNetCode = int //类型别名

const myNetCode NetCode = 1
const myAliasNetCode AliasNetCode = 1

//区别1：自定义类型可以绑方法，类型别名不能绑方法，也技术不能 func () funcName ()()

func main() {

	//区别2：类型别名本质上还是基础类型
	fmt.Printf("%v,%T\n", myNetCode, myNetCode)
	fmt.Printf("%v,%T\n", myAliasNetCode, myAliasNetCode)

	var judge int = 1

	//区别3:自定义类型肯定不能跟基础类型比较，但类型别名可以
	fmt.Println(judge == myAliasNetCode)

}
