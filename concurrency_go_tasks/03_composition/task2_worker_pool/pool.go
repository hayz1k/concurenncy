package pool

import (
	"sync"
	"time"
)

func RunPool(jobs []int, workers int) int {
	jobCh := make(chan int)
	resultCh := make(chan int, len(jobs))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobCh {
				time.Sleep(10 * time.Millisecond)
				resultCh <- j
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for _, j := range jobs {
		jobCh <- j
	}
	close(jobCh)

	sum := 0
	for r := range resultCh {
		sum += r
	}
	return sum
}
