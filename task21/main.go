package main

// Movable - интерфейс, который мы хотим реализовать.
type Movable interface {
	Move(x, y, z float64)
}

// Duck - структура, которая должна реализовывать наш интерфейс,
// но не реализует, поскольку структура не имеет метода Move, хоть
// и имеет схожий метод Swim.
type Duck struct{}

func (Duck) Swim(x, y, z float64) {}

// DuckAdapter - собственно адаптер, который позволит нам реализовать
// интерфейс Movable для структуры Duck.
type DuckAdapter struct {
	Duck
}

func NewDuckAdapter(duck Duck) *DuckAdapter {
	return &DuckAdapter{Duck: duck}
}

// Move будет использовать доступный в структуре Duck метод Swim.
func (d DuckAdapter) Move(x, y, z float64) {
	d.Swim(x, y, z)
}

// Consumer - некоторая функция, которая будет работать только со
// значениями типа Movable.
func Consumer(movable Movable) {
	movable.Move(1, 2, 3)
}

func main() {
	duck := &Duck{}
	adapter := NewDuckAdapter(*duck)
	Consumer(adapter)
}
