package scheduler

import "time"

func Every(d time.Duration, f func()) (stop func()) {
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				f()
			case <-done:
				return
			}
		}
	}()
	return func() { close(done) }
}
