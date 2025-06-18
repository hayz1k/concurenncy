package generator

import "context"

// Generate возвращает канал, из которого можно читать возрастающие числа,
// начиная с нуля. Генерация прекращается при отмене ctx.
func Generate(ctx context.Context) <-chan int {
	// TODO: реализовать генератор чисел с учётом отмены
	return nil
}
