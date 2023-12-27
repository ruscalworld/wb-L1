package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

// Функция worker получает числа из канала numbers и выводит их, используя Printer.
func worker(numbers chan float64, result *atomic.Int32, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for request := range numbers {
		result.Add(int32(math.Pow(request, 2)))
	}
}

// Функция processNumbers отправляет числа из переданного в неё слайса numbers
// в канал numbers.
func processNumbers(numbers chan float64, task []float64) {
	for _, number := range task {
		numbers <- number
	}
}

// Функция makeWorkers создаёт горутины, которые будут прибавлять квадраты чисел из канала numbers к
// result. Возвращает WaitGroup, который можно использовать для ожидания их завершения.
func makeWorkers(numbers chan float64, result *atomic.Int32, amount int) *sync.WaitGroup {
	// Создаём WaitGroup, с помощью которого будем ждать завершения всех горутин.
	waitGroup := &sync.WaitGroup{}

	// Создаём нужное количество горутин.
	for i := 0; i < amount; i++ {
		waitGroup.Add(1)
		go worker(numbers, result, waitGroup)
	}

	return waitGroup
}

func main() {
	// Создаём канал, через который будем отправлять числа в горутины.
	numbers := make(chan float64)

	// Создаём переменную, в которой будем хранить результат.
	// Используем атомики, чтобы избежать проблем с конкурентным прибавлением чисел к результату.
	result := &atomic.Int32{}

	// Создаём 2 горутины для обработки чисел.
	waitGroup := makeWorkers(numbers, result, 2)

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

	// Выводим результат.
	fmt.Println(result.Load())
}
