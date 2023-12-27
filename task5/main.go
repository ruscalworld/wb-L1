package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// Printer - структура, предотвращающая одновременный вывод
// из нескольких горутин.
type Printer struct {
	lock sync.Mutex
}

// Print выводит переданные данные в стандартный вывод
func (p *Printer) Print(data ...any) {
	// Используем мьютексы, чтобы предотвратить конкурентную запись в поток вывода
	p.lock.Lock()
	defer p.lock.Unlock()

	// Выводим данные
	fmt.Println(data...)
}

// Функция worker получает данные из канала data и выводит их с помощью Printer.
func worker(ctx context.Context, data chan string, printer *Printer, waitGroup *sync.WaitGroup) {
	// При выходе из функции, который означает завершение работы горутины, сообщаем о том, что мы закончили.
	defer waitGroup.Done()

	// В бесконечном цикле получаем данные их канала, либо завершаем работу, если получили
	// сигнал о завершении через контекст.
	for {
		select {
		case d := <-data:
			// Выводим данные из канала.
			printer.Print(d)
		case <-ctx.Done():
			// Выводим сообщение об окончании работы.
			printer.Print("Shutting down worker")

			// Выходим из функции, тем самым завершая горутину.
			return
		}
	}
}

// Функция writer бесконечно записывает случайные числа в канал data.
func writer(ctx context.Context, data chan string, printer *Printer, waitGroup *sync.WaitGroup) {
	// При выходе из функции, который означает завершение работы горутины, сообщаем о том, что мы закончили.
	defer waitGroup.Done()

	// В бесконечном цикле либо отправляем случайное число в канал, либо завершаем работу, если получили
	// сигнал о завершении через контекст.
	for {
		select {
		default:
			// Отправляем случайное число в канал.
			data <- strconv.Itoa(rand.Int())

			// Ждём 100 мс.
			time.Sleep(100 * time.Millisecond)
		case <-ctx.Done():
			// Выводим сообщение об окончании работы.
			printer.Print("Shutting down writer")

			// Выходим из функции, тем самым завершая горутину.
			return
		}
	}
}

func main() {
	// Создаём Printer, который будет использоваться для вывода данных.
	printer := &Printer{}

	// Создаём WaitGroup, с помощью которого будем ждать завершения всех горутин.
	waitGroup := &sync.WaitGroup{}

	// Создаём канал, через который будем отправлять данные в горутины.
	data := make(chan string)

	// Создаём контекст с таймаутом в 5 секунд. Таким образом, через 5 секунд
	// горутины, использующие этот контекст, завершатся.
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	// Запускаем горутину, которая будет записывать случайные числа в канал.
	waitGroup.Add(1)
	go writer(ctx, data, printer, waitGroup)

	// Запускаем воркер, который будет выводить получаемые данные.
	waitGroup.Add(1)
	go worker(ctx, data, printer, waitGroup)

	// Ждём завершения всех горутин.
	waitGroup.Wait()
}
