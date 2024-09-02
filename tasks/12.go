//	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для
//	нее собственное множество.

// Комментарий автора
// Теперь, когда мы выяснили что есть множество мы можем без особых
// усилий и проверок выполнить данное задание
package tasks

import (
	"fmt"
	"sort"
)

// Используем мапу структур (в пустой структуре 0 размер памяти)
// И будем просто создавать string ключи
func Task12_1() {

	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	var m = make(map[string]struct{})
	var set []string

	//Запись ключей map
	for i := 0; i < 5; i++ {
		m[arr[i]] = struct{}{}
	}

	//Запись слов в массив
	for key := range m {
		set = append(set, key)
	}

	fmt.Println(set)

}

// Можем и сделать и через сортировку. Естественно, это куда менее эффективно
func Task12_2() {
	arr := make([]string, 0)
	arr = append(arr, "cat", "cat", "dog", "cat", "tree")
	sort.Strings(arr)
	for i := 1; i < len(arr); {
		if arr[i] == arr[i-1] {
			arr = append(arr[:i-1], arr[i:]...)
		} else {
			i++
		}
	}
	fmt.Println(arr)
}
