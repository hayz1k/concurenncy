package debounce

import "time"

func Debounce(d time.Duration, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		var (
			timer *time.Timer
			last  int
			ok    bool
		)
		for {
			var c <-chan time.Time
			if timer != nil {
				c = timer.C
			}
			select {
			case v, more := <-in:
				if !more {
					if ok {
						if timer != nil {
							timer.Stop()
						}
						out <- last
					}
					return
				}
				last = v
				ok = true
				if timer != nil {
					if !timer.Stop() {
						<-timer.C
					}
				}
				timer = time.NewTimer(d)
			case <-c:
				out <- last
				ok = false
				timer = nil
			}
		}
	}()
	return out
}
