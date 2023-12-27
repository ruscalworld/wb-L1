package main

import (
	"fmt"
	"math"
	"sync"
)

// Printer - структура, предотвращающая одновременный вывод
// из нескольких горутин.
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

// Функция worker получает числа из канала numbers, вычисляет их квадраты и выводит, используя Printer.
func worker(numbers chan float64, printer *Printer, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for request := range numbers {
		result := math.Pow(request, 2)
		printer.PrintResult(request, result)
	}
}

// Функция processNumbers отправляет числа из переданного в неё слайса numbers
// в канал c.
func processNumbers(numbers chan float64, task []float64) {
	for _, number := range task {
		numbers <- number
	}
}

// Функция makeWorkers создаёт горутины, которые будут возводить числа из канала c в квадрат,
// и возвращает WaitGroup, который можно использовать для ожидания их завершения.
func makeWorkers(numbers chan float64, amount int) *sync.WaitGroup {
	// Создаём Printer, который будет использоваться для вывода результата.
	printer := &Printer{}

	// Создаём WaitGroup, с помощью которого будем ждать завершения всех горутин.
	waitGroup := &sync.WaitGroup{}

	// Создаём нужное количество горутин.
	for i := 0; i < amount; i++ {
		waitGroup.Add(1)
		go worker(numbers, printer, waitGroup)
	}

	return waitGroup
}

func main() {
	// Создаём канал, через который будем отправлять числа в горутины.
	numbers := make(chan float64)

	// Создаём 2 горутины для обработки чисел.
	waitGroup := makeWorkers(numbers, 2)

	// Создаём слайс с числами из задания.
	task := []float64{2, 4, 6, 8, 10}

	// Отправляем числа из слайса в горутины.
	processNumbers(numbers, task)

	// Когда все числа отправлены, закрываем канал.
	// После закрытия канала for внутри функции worker завершится.
	close(numbers)

	// После завершения цикла в горутине вызовется функиця Done() для
	// WaitGroup, что позволит нам отследить завершение горутины.
	// Когда все горутины завершатся, Wait разблокирует текущую горутину.
	waitGroup.Wait()
}
