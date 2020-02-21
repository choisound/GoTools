package test

import (
	"fmt"
	"time"
)

//生产者
func Producer(queue chan<- int) {
	for i := 0; i < 10; i++ {
		queue <- i
		fmt.Printf("生产了%d\n", i)
		time.Sleep(1 * time.Second)
	}
}

//消费者
func Consumer(queue <-chan int) {
	for i := 0; i < 10; i++ {
		i := <-queue
		fmt.Printf("消费了%d\n", i)
	}
}
