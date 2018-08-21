package runnner

import (
	"github.com/gpmgo/gopm/modules/log"
	"daker.wang/Azen/Go-execise/Streaming/dispatcher/ops"
	"sync"
)

type TaskEvent int

const (
	TaskEventNormal = iota // 本次生产结束，正常执行下次生产
	TaskEventClose         // 结束所有生产
	TaskEventError         // 结束所有生产 & 遇到错误
)

type TaskResult struct {
	Event TaskEvent
	Err map[string]error
}

func VideoClearDispatcher(data chan string) *TaskResult {

	ids, err := ops.GetPaddingVIDs(3); if err != nil {
		log.Error(err.Error())
		return &TaskResult{Event: TaskEventError, Err:map[string]error{"sp":err}}
	}

	if len(ids)  <= 0 {
		log.Warn("🎉 数据已经处理完毕")
		return &TaskResult{Event: TaskEventClose}
	}

	//  分发指定数量的数据到仓库中
	for _, id := range ids {
		log.Warn("Dispatch - %s", id)
		data <- id
	}

	return &TaskResult{Event:TaskEventNormal}
}


func VideoClearExecutor(data chan string) *TaskResult {

	errMap := sync.Map{}

	forloop:
		for {
			select {
			case d := <- data:
				log.Warn("Executor - %s", d)
				go func(item string) {
					err := ops.DeleteVideo(item); if err != nil {
						errMap.Store(item, err)
						log.Error("delete error id：%s, err: %v", item, err.Error())
					}
					log.Warn("🎉 %s的生产任务结束了", item)

				}(d)
			default:
				break forloop
			}
		}

	errs := pickUpSyncMap(errMap)
	if errs == nil {
		return &TaskResult{Event:TaskEventNormal}
	} else {
		return &TaskResult{Event:TaskEventError, Err:errs}
	}
}

func pickUpSyncMap(syncMap sync.Map) map[string]error {
	var errResult = map[string]error{}
	syncMap.Range(func(key, value interface{}) bool {
		errResult[key.(string)] = value.(error)
		return true
	})
	return errResult
}
