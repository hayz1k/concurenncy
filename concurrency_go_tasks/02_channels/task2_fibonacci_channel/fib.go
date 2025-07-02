package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 0, 1
		if n > 0 {
			ch <- a
		}
		for i := 1; i < n; i++ {
			a, b = b, a+b
			ch <- a
		}
	}()

	// TODO: отправить последовательность Фибоначчи в канал
	return ch
}
