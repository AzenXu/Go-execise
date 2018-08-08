package limiter

import (
	"github.com/gpmgo/gopm/modules/log"
)

type Limiter struct {
	ConcurrentCount int
	TokenBucket chan int
}

func New(cc int) (limiter *Limiter) {

	tb := make(chan int, cc)

	l := &Limiter{ConcurrentCount:cc, TokenBucket:tb}

	return l
}

func(l *Limiter) GetToken() (ok bool) {
	if len(l.TokenBucket) >= l.ConcurrentCount {
		log.Warn("限流机制启动")
		return false
	}

	log.Warn("--- ❌ Dispatch a token ---")
	l.TokenBucket <- 1

	return true
}

func(l *Limiter) ReleaseToken() {
	if len(l.TokenBucket) <= 0 {
		return
	}

	<- l.TokenBucket
	log.Warn("--- ✅ Bucket got a empty token ---")
}