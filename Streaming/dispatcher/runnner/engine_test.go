package runnner

import (
	"testing"
	"github.com/gpmgo/gopm/modules/log"
	"time"
	"strconv"
)

func TestEngine_StartAll(t *testing.T) {

	indexd := 0
	indexe := 0
	en := New(3, func(data chan string) error {
		log.Warn("🎉 第%d次生产", indexd)
		indexd++
		for i := 0; i < 3; i++ {
			go func(idx int) { // 使用go程，是为了模拟生产者消费者模式中的，生产者并发模型
				log.Warn("Dispatch - %d", idx)
				data <- strconv.Itoa(idx)
			}(i)
		}
		time.Sleep(1 * time.Second)
		return nil
	}, func(data chan string) error {
		log.Warn("⌲ 第%d次消费", indexe)
		indexe++
		forloop:
			for {
				select {
				case d := <- data:
					log.Warn("Executor - %s", d)
				default:
					break forloop
				}
			}

		time.Sleep(1 * time.Second)
		return nil
	})

	go en.StartAll()

	time.Sleep(10 * time.Second)
}