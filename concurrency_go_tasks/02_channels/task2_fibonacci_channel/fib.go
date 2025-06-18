package fibonacci

// Fib возвращает канал, из которого можно читать первые n чисел Фибоначчи.
func Fib(n int) <-chan int {
	ch := make(chan int)
	close(ch)
	// TODO: отправить последовательность Фибоначчи в канал
	return ch
}
