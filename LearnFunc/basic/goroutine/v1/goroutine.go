package main

import (
	"log"
	"math/rand"
	"time"
)

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 睡眠片刻（随机0到2.5秒）
	}
}

func main() {
	//rand.Seed(time.Now().UnixNano()) // Go 1.20之前需要
	//使用log打印是因为log是线程安全的有并发同步
	log.SetFlags(0)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	time.Sleep(10 * time.Second)
}
