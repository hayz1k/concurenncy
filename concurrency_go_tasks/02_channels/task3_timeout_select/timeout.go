package timeout

import "context"

// Work выполняет длительную задачу и возвращает ошибку,
// если она заняла больше 100 мс или контекст был отменён.
func Work(ctx context.Context) error {
	// TODO: реализовать через select и time.After
	return nil
}
