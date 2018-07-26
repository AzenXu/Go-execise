package scheduler

import (
	"daker.wang/Azen/Go-execise/Demo/Spider/Engine"
	"log"
)

type QueueScheduler struct {

	workerInputChannel chan chan engine.Request
	taskInputChannel chan engine.Request

	tasks []engine.Request
	workers []chan engine.Request
}

func (scheduler *QueueScheduler) GetRequsetChannel() (requestChannel chan engine.Request) {
	return make(chan engine.Request)
}

func (scheduler *QueueScheduler) Submit(request engine.Request) {
	if len(request.URL) <= 0 {
		return
	}
	scheduler.taskInputChannel <- request
}

func (scheduler *QueueScheduler) SubmitWorker(worker chan engine.Request) {
	scheduler.workerInputChannel <- worker
}

func (scheduler *QueueScheduler) Run() {

	scheduler.workerInputChannel = make(chan chan engine.Request)
	scheduler.taskInputChannel = make(chan engine.Request)

	go func() {
		for {

			var nextTask engine.Request
			var nextWorker chan engine.Request

			if len(scheduler.tasks) > 0 && len(scheduler.workers) > 0 {
				nextTask = scheduler.tasks[0]
				nextWorker = scheduler.workers[0]
			}

			select {
			case worker := <- scheduler.workerInputChannel:
				scheduler.workers = append(scheduler.workers, worker)
			case task := <- scheduler.taskInputChannel:
				scheduler.tasks = append(scheduler.tasks, task)
			case nextWorker <- nextTask:
				scheduler.tasks = scheduler.tasks[1:]
				scheduler.workers = scheduler.workers[1:]
				log.Println("配对成功~~~")
			}
		}
	}()
}

/*
待实现func
*/
func (scheduler *QueueScheduler) Remove(request engine.Request) {
	// 从队列里找到request然后移除掉
	panic("待实现")
}

func (scheduler *QueueScheduler) StopWorker(worker chan engine.Request) {
	/*
	有两个思路控制task：
	1. 传特殊的request（控制request）进来，worker检测到这种request，就在for循环中return掉
	2. 在插一根管儿，这根管专门发控制信号 - 停止、快一点、慢一点、中断任务处理
	*/
	panic("待实现")
}