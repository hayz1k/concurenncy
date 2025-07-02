package producerconsumer

import (
	"fmt"
	"io"
	"sync"
)

// Run запускает продюсера, который отправляет числа от 1 до 10, и консюмера,
// который выводит их в writer. Используйте небуферизованный канал и ожидание
// завершения горутин.
func Run(w io.Writer) {
	// TODO: реализовать продюсер и консюмер
	producer := make(chan int)
	consumer := make(chan int)
	var wg sync.WaitGroup
	wg.Add(3)

	// producer
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			producer <- i
		}
		close(producer)
	}()

	// consumer
	go func() {
		defer wg.Done()
		for value := range producer {
			consumer <- value
		}
		close(consumer)
	}()

	go func() {
		defer wg.Done()
		for value := range consumer {
			fmt.Fprintln(w, value)
		}
	}()

	wg.Wait()
}
