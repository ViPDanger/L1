// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской
// структуры Human (аналог наследования).
package tasks

import "fmt"

// родительская структура Human
type human struct {
	name string
}

// дочерняя структура с встраиванием методов
type action struct {
	name string
	human
}

// функция human
func (h human) doNothing() {
	fmt.Println("Nah, i'm too tired, because i'm", h.name)
}

func Task1_1() {
	act := action{
		name: "Action",
		human: human{
			name: "Human",
		},
	}
	act.doNothing()
}
