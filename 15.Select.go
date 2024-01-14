package main

import (
	"fmt"
	"sync"
	"time"
)

func shoppingAndBuyAndRecord(name string, money int, waitGroup *sync.WaitGroup) {
	fmt.Println(name, "在购物\n")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "购物结束\n")

	moneyChannel_1 <- money
	nameChannel_1 <- name
	waitGroup.Done()
}

// 声明并初始化为一个 长度为0的channel
// channel是一个上了锁的实例
var moneyChannel_1 = make(chan int)
var nameChannel_1 = make(chan string)
var doneChannel = make(chan struct{})

// 注：主线程结束，协程就会结束,所以必须要让主线程等待所有协程完成

func main() {
	startTime := time.Now()
	var waitGroup sync.WaitGroup
	fmt.Println("-------所有人购物开始--------")

	waitGroup.Add(2)
	go shoppingAndBuyAndRecord("xie", 4, &waitGroup)
	go shoppingAndBuyAndRecord("yang", 1, &waitGroup)

	go func() {
		defer close(moneyChannel_1)
		defer close(nameChannel_1)
		defer close(doneChannel)
		waitGroup.Wait()

	}()

	var moneyList []int
	var nameList []string
	var event = func() {
		for {
			select {
			case money := <-moneyChannel_1:
				moneyList = append(moneyList, money)
				println("接到数据money")
			case name := <-nameChannel_1:
				nameList = append(nameList, name)
				println("接到数据name")
			case <-doneChannel:
				println("--------doneChannel-------")
				return
			default:
				println("未接到数据")
			}
		}
	}
	event()
	fmt.Println("-------所有人购物结束--------")
	duringTime := time.Since(startTime)
	fmt.Println("耗时: ", duringTime)
	fmt.Println("总共花了钱: ", moneyList)
	fmt.Println("购物的人有", nameList)

}
