package limiter

import "time"

type Limiter struct {
	tokens chan struct{}
}

func NewLimiter() *Limiter {
	l := &Limiter{tokens: make(chan struct{}, 5)}
	for i := 0; i < 5; i++ {
		l.tokens <- struct{}{}
	}
	go func() {
		ticker := time.NewTicker(time.Second / 5)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case l.tokens <- struct{}{}:
			default:
			}
		}
	}()
	return l
}

func (l *Limiter) Allow() bool {
	select {
	case <-l.tokens:
		return true
	default:
		return false
	}
}
