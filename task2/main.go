package main

import (
	"fmt"
	"math"
	"sync"
)

// Printer - структура, отвечающая за вывод, предотвращающая одновременный вывод
// из нескольких потоков.
type Printer struct {
	lock sync.Mutex
}

// PrintResult выводит исходное число и его квадрат в стандартный вывод
func (p *Printer) PrintResult(req, res float64) {
	// Используем мьютексы, чтобы предотвратить конкурентную запись в поток вывода
	p.lock.Lock()
	defer p.lock.Unlock()

	// Выводим результат
	fmt.Printf("%.0f^2 = %.0f\n", req, res)
}

// Функция worker получает числа из канала numbers и выводит их, используя Printer.
func worker(numbers chan float64, printer *Printer, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for request := range numbers {
		result := math.Pow(request, 2)
		printer.PrintResult(request, result)
	}
}

func main() {
	// Создаём канал, через который будем отправлять числа в горутины.
	numbers := make(chan float64)

	// Создаём Printer, который будет использоваться для вывода результата.
	printer := &Printer{}

	// Создаём WaitGroup, с помощью которого будем ждать завершения всех горутин.
	w := &sync.WaitGroup{}

	// Создаём 5 горутин.
	for i := 0; i < 5; i++ {
		w.Add(1)
		go worker(numbers, printer, w)
	}

	// Отправляем числа в созданные горутины с помощью канала.
	for i := 0.0; i < 20; i++ {
		numbers <- i
	}

	// Когда все числа отправлены, закрываем канал.
	// После закрытия канала for внутри функции worker завершится.
	close(numbers)

	// После завершения цикла в горутине вызовется функиця Done() для
	// WaitGroup, что позволит нам отследить завершение горутины.
	// Когда все горутины завершатся, Wait разблокирует текущую горутину.
	w.Wait()
}
