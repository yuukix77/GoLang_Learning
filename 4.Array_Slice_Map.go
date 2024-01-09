package main

import (
	"fmt"
	"sort"
)

func main() {

	//数组, 不能增删,只能改查
	var nameList = [3]string{"小红", "小黄", "小绿"}
	fmt.Println(nameList)
	fmt.Println(nameList[0])
	fmt.Println(nameList[len(nameList)-1])

	//切片或列表，可以增删改查
	var colorList []string
	colorList = append(colorList, "blue")
	colorList = append(colorList, "red")
	fmt.Println(colorList)

	var nullList []string //定义为空
	fmt.Println(nullList)
	fmt.Println(nullList == nil) //nil是go语言的null

	var ageList = make([]int, 3) //用make初始化, 创建指定容量的slice, 用0来填充未赋值的位置
	ageList = append(ageList, 1)
	fmt.Println(ageList)

	//切片切数组
	var arrayTest = [6]int{1, 2, 3, 4, 5, 6}
	slice1 := arrayTest[:]
	slice2 := arrayTest[:2]
	fmt.Println(slice1)
	fmt.Println(slice2)

	//排序操作
	var arrayTest2 = []int{7, 3, 3, 8, 5, 6}
	sort.Ints(arrayTest2) //升序
	fmt.Println(arrayTest2)
	sort.Sort(sort.Reverse(sort.IntSlice(arrayTest2))) //降序
	fmt.Println(arrayTest2)

}
