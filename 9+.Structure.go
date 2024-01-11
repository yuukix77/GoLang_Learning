package main

import "fmt"

type User struct {
	name string
	age  int
}

func (this *User) setName(name string) {
	this.name = name

}

func main() {
	var user = User{name: "John", age: 11}
	fmt.Println(user)
	user.setName("yuki")
	fmt.Println(user)
}
