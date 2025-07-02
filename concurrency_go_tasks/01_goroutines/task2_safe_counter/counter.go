package counter

import "sync"

// Counter хранит целое значение и мьютекс для безопасного доступа.
type Counter struct {
	mu sync.Mutex
	v  int
}

// Inc увеличивает счётчик на 1 с защитой от гонок.
func (c *Counter) Inc() {
	// TODO: реализовать инкремент с использованием мьютекса
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v++
}

// Value возвращает текущее значение счётчика безопасно для гонок.
func (c *Counter) Value() int {
	// TODO: вернуть значение с учётом блокировки
	c.mu.Lock()
	defer c.mu.Unlock()
	value := c.v
	return value
}
