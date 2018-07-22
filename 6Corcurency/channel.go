package main

import (
	"fmt"
	"time"
	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	//chanRight()
	//channelAsParam()
	//channelMuchMoreWorker()
	//bufferedChannelDemo()
	//closeChannel()
	workerCloseRangeDemo()
}

func createWorker(id int) chan<- int {
	channel := make(chan int)
	go worker(id, channel)
	return channel
}

func worker(id int, channel chan int) {
	log.Warn("%d - 已经就位~", id)

	for {
		content, ok := <- channel
		if !ok {
			log.Error("那个...被Close了呢...")
			break
		}
		log.Warn("id是%d的协程，收到了：%d", id, content)
		time.Sleep(100 * time.Millisecond)
	}
}

func workerCloseRangeDemo() {
	channel := make(chan int)
	go workerCloseJudgementForRange(channel)

	channel <- 100

	close(channel)
	time.Sleep(1 * time.Second)
}

func workerCloseJudgementForRange(channel chan int) {
	for n := range channel {
		log.Warn("收到了：%d", n)
	}
}

func closeChannel() {
	c := make(chan int)
	go worker(0 , c)
	c <- 1
	c <- 2
	c <- 3
	close(c)
	time.Sleep(1 * time.Second)
}

func bufferedChannelDemo() {
	channel := make(chan int, 3)
	go worker(0, channel)
	channel <- 10
	log.Warn("10发送完毕 - 你看我你看我~O(∩_∩)O~")
	channel <- 20
	log.Warn("20发送完毕 - 你看我你看我~O(∩_∩)O~")
	channel <- 30
	log.Warn("30发送完毕 - 你看我你看我~O(∩_∩)O~")
	channel <- 40
	log.Warn("40发送完毕 - 你看我你看我~O(∩_∩)O~")
	time.Sleep(time.Second)
}

func channelMuchMoreWorker() {
	var channels [10]chan int

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)

		go worker(i, channels[i])
	}

	time.Sleep(1 * time.Second)

	log.Warn("------准备发送信号了------")

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(1 * time.Second)
}

func channelAsParam() {
	c := make(chan int)

	go worker(0, c)

	c <- 8
	time.Sleep(1 * time.Second)
}

func chanRight() {
	c := make(chan int)

	go func() {
		fmt.Println("准备发送了")
		time.Sleep(10 * time.Second)
		c <- 1
		fmt.Println("发送完毕了")
	}()

	fmt.Println("准备接收了")
	n := <- c
	time.Sleep(time.Second)
	fmt.Println("接收到了：", n)
}

func chanDemoError() {
	c := make(chan int)
	c <- 1
	c <- 2
	n := <- c
	fmt.Println(n)
}
