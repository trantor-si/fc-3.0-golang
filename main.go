package main

import (
	"fmt"
	"time"
)

func worker1(done chan string) {
	fmt.Println("worker1 started")
	time.Sleep(time.Second * 2)
	fmt.Println("worker1 id done")
	done <- "Worker1 f*"
}

func worker2(done chan string) {
	fmt.Println("worker2 started")
	time.Sleep(time.Second * 2)
	fmt.Println("worker2 id done")
	done <- "Worker2 f*"
}

func main() {
	done := make(chan string)

	go worker1(done)
	go worker2(done)

	fmt.Println(<-done)
	time.Sleep(time.Second * 10)
}
