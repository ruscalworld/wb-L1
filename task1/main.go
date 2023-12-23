package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

// IsAdult - метод для структуры Human.
func (h *Human) IsAdult() bool {
	return h.Age >= 18
}

// Action - структура, содержащая вложенную структуру Human.
// Таким образом Action "наследует" от Human все поля и методы.
type Action struct {
	Human
	SomeAnotherField any
}

// GetDescription - метод для Action, который использует полученный от структуры Human
// метод IsAdult.
func (a *Action) GetDescription() string {
	if a.IsAdult() {
		return fmt.Sprintf("%s - совершеннолетний", a.Name)
	} else {
		return fmt.Sprintf("%s - несовершеннолетний", a.Name)
	}
}

func main() {
	// Для проверки создадим массив из двух Action, которые будут содержать в себе
	// Human с возрастом 12 и 20.
	actions := []Action{
		{
			Human: Human{
				Age:  12,
				Name: "Петя",
			},
		},
		{
			Human: Human{
				Age:  20,
				Name: "Василий",
			},
		},
	}

	for _, action := range actions {
		// Здесь вызовем метод GetDescription из Action, который в свою очередь вызовет
		// метод IsAdult, полученный из вложенного Human.
		fmt.Println(action.GetDescription())
	}
}
