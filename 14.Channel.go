package main

import (
	"fmt"
	"sync"
	"time"
)

func shoppingAndBuy(name string, money int, wait *sync.WaitGroup) {
	fmt.Println(name, "在购物\n")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "购物结束\n")

	moneyChannel <- money

	wait.Done()
}

// 声明并初始化为一个 长度为0的channel
// channel是一个上了锁的实例
var moneyChannel = make(chan int)

// 注：主线程结束，协程就会结束,所以必须要让主线程等待所有协程完成

func main() {
	startTime := time.Now()
	moneySum := 0
	var wait sync.WaitGroup
	fmt.Println("-------所有人购物开始--------")

	wait.Add(2)
	go shoppingAndBuy("xie", 4, &wait)
	go shoppingAndBuy("yang", 1, &wait)

	go func() {
		wait.Wait()
		close(moneyChannel)
	}()

	for money := range moneyChannel {
		moneySum += money
	}

	wait.Wait()

	fmt.Println("-------所有人购物结束--------")
	duringTime := time.Since(startTime)

	fmt.Println("耗时: ", duringTime)
	fmt.Println("总共花了钱: ", moneySum)

}
