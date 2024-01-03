package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter - общий интерфейс для структуры, требуемой по заданию.
type Counter interface {
	Inc()
	Value() int64
}

// AtomicCounter использует atomic.Int64, который гарантирует безопасное
// конкурентное использование, для хранения текущего значения счётчика.
type AtomicCounter struct {
	value atomic.Int64
}

func (a *AtomicCounter) Inc() {
	// Используем метод Add, который прибавляет нужное нам значение (1)
	// к текущему значению внутри atomic.
	a.value.Add(1)
}

func (a *AtomicCounter) Value() int64 {
	return a.value.Load()
}

// MutexCounter использует sync.Mutex для обеспечения безопасного конкурентного
// доступа к значению, которое, в свою очередь, хранится в чистом виде.
type MutexCounter struct {
	value int64
	lock  sync.Mutex
}

func (m *MutexCounter) Inc() {
	// Блокируем мьютекс, чтобы избежать конкурентного взаимодействия с полем value.
	m.lock.Lock()
	defer m.lock.Unlock()

	m.value++
}

func (m *MutexCounter) Value() int64 {
	return m.value
}

// Функция, которая будет заданное в amount количество раз делать инкремент
// переданного в неё Counter.
func incrementor(amount int, counter Counter, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for i := 0; i < amount; i++ {
		counter.Inc()
	}
}

func main() {
	waitGroup := &sync.WaitGroup{}
	//counter := &MutexCounter{}
	counter := &AtomicCounter{}

	// Запустим 5 горутин, которые увеличат значение нашего счётчика на
	// 50 + 40 + 30 + 20 + 10 = 150.
	waitGroup.Add(5)
	go incrementor(50, counter, waitGroup)
	go incrementor(40, counter, waitGroup)
	go incrementor(30, counter, waitGroup)
	go incrementor(20, counter, waitGroup)
	go incrementor(10, counter, waitGroup)

	waitGroup.Wait()
	fmt.Println(counter.Value()) // 150
}
