package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Println(name, "在购物\n")
	time.Sleep(1 * time.Second)
	fmt.Println(name, "购物结束\n")
	wait.Done()
}

// 注：主线程结束，协程就会结束,所以必须要让主线程等待所有协程完成

func main() {
	startTime := time.Now()
	var wait sync.WaitGroup
	fmt.Println("-------所有人购物开始--------")

	wait.Add(2)
	go shopping("xie", &wait)
	go shopping("yang", &wait)
	wait.Wait()

	fmt.Println("-------所有人购物结束--------")
	duringTime := time.Since(startTime)

	fmt.Println("耗时: ", duringTime)

}
