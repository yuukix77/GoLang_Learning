package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Name     string `json:"name"`
	Id       int    `json:"id,omitempty" //忽略空值`
	Password string `json:"password"` //json:"-" 表示忽略
}

func main() {
	var user = UserInfo{Name: "Mystic", Id: 2, Password: "sjasghd"}
	byteData, _ := json.Marshal(user)
	fmt.Println(byteData)
	fmt.Println(string(byteData))

}
