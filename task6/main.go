package main

import (
	"context"
	"fmt"
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

// Через контекст.
// Внутри горутины выполняется бесконечный цикл. Если приходит команда о завершении через канал,
// выполняется соответствующая ветка в select, завершающая выполнение бесконечного цикла.
func method1(ctx context.Context, printer *Printer, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	i := 0

	for {
		select {
		default:
			i += 1
		case <-ctx.Done():
			printer.Print("Горутина 1 завершена")
			return
		}
	}
}

// Через канал.
// То же самое, что и с контекстами, только сами создаём канал для оповещения о завершении работы.
// Внутри горутины выполняется бесконечный цикл. Если приходит команда о завершении через канал,
// выполняется соответствующая ветка в select, завершающая выполнение бесконечного цикла.
func method2(done chan bool, printer *Printer, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	i := 0

	for {
		select {
		default:
			i += 1
		case <-done:
			printer.Print("Горутина 2 завершена")
			return
		}
	}
}

func main() {
	waitGroup := &sync.WaitGroup{}
	printer := &Printer{}

	// Метод 1.
	// Создаём контекст с возможностью отмены, после чего можем использовать CancelFunc
	// для завершения работы горутины.
	ctx, cancel1 := context.WithCancel(context.Background())
	waitGroup.Add(1)
	go method1(ctx, printer, waitGroup)

	// Метод 2.
	// Создаём канал, через который будем сообщать горутине о завершении, после чего
	// можем отправить в него значение для завершения работы горутины.
	cancel2 := make(chan bool)
	waitGroup.Add(1)
	go method2(cancel2, printer, waitGroup)

	// Дадим горутинам время запуститься таким костыльным способом.
	time.Sleep(time.Second)

	cancel1()
	cancel2 <- true
	waitGroup.Wait()
}
