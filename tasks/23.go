// Удалить i-ый элемент из слайса.
package tasks

import (
	"fmt"
	"log"
)

// Красивая функция delete для любого слайса
func delete[Type any](idForDelete int, slice []Type) []Type {
	if idForDelete > len(slice) || idForDelete < 0 {
		log.Fatalln("Invalid IdForDelete")
	}
	return append(slice[:idForDelete], slice[idForDelete+1:]...)
}

func Task23_1() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	id := 3
	fmt.Println(delete(id, slice))
}

// Можно написать и без неё. Просто пропишем append для конкретного слайса
func Task23_2() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	id := 3
	if id > len(slice) || id < 0 {
		log.Fatalln("Invalid IdForDelete")
	}
	fmt.Println(append(slice[:id], slice[id+1:]...))
}
