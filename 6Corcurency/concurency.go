package main

import (
	"fmt"
	"strconv"
	"time"
	"runtime"
)

func main() {
	//concurrencyTest()
	//concurrencyNoGiveup()
	//concurrencyIndexTest()
	concurrencyGiveup()
}

func concurrencyGiveup() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(index int) {
			for {
				a[index]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(a)
}

func concurrencyIndexTest() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() {
			for {
				a[i]++
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(a)
}

//  不交出控制权
func concurrencyNoGiveup() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(index int) {
			for {
				a[index]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
	fmt.Println(a)
}

func concurrencyTest() {
	for i := 0; i < 1000; i++ {
		go func() {
			for {
				fmt.Println("好饿好饿好饿我真的好饿...我是：" + strconv.Itoa(i))
			}
		}()
	}

	time.Sleep(2 * time.Millisecond)
}