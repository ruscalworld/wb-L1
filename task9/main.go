package main

import (
	"fmt"
	"math"
	"sync"
)

// На втором этапе получаем числа из канала и выводим их.
func secondStep(numbers chan float64, wg *sync.WaitGroup) {
	defer wg.Done()

	for number := range numbers {
		fmt.Println(number)
	}
}

// На первом этапе получаем числа из первого канала и записываем их квадраты во второй канал.
func firstStep(src, dst chan float64, wg *sync.WaitGroup) {
	defer wg.Done()

	// Закрываем второй канал, когда все данные в него уже записаны.
	defer close(dst)

	for number := range src {
		dst <- math.Pow(number, 2)
	}
}

func main() {
	// Создаём каналы и WaitGroup.
	wg := &sync.WaitGroup{}
	src := make(chan float64)
	dst := make(chan float64)

	// Запускаем горутины, которые будут работать с каналами.
	wg.Add(2)
	go firstStep(src, dst, wg)
	go secondStep(dst, wg)

	// Записываем числа от 0 до 20 в первый канал.
	for i := 0.0; i < 20; i++ {
		src <- i
	}

	// Закрываем первый канал, когда все данные в него уже записаны.
	close(src)

	// Ждём завершения всех горутин.
	wg.Wait()
}
