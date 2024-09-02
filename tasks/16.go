// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка
package tasks

import "fmt"

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	// wall - стена. ID обьекта, после которого начинается неотсортированный массив
	wall := 0
	// pivot - Число, относительно которого сортируется массив
	pivot := a[len(a)-1]
	// Цикл сортировки
	for current := 0; current < len(a)-1; current++ {
		// Кидаем обьект в сортированное и сдвигаем wall на один
		if a[current] <= pivot {
			a[wall], a[current] = a[current], a[wall]
			wall++
		}
	}
	// Меняем местами wall и последний элемент массива, после чего запускаем quicksort на каждую из них
	a[wall], a[len(a)-1] = a[len(a)-1], a[wall]
	quickSort(a[0:wall])
	quickSort(a[wall:])
	return a
}

func Task16_1() {
	// Пример сортируемого массива
	arr := make([]int, 0)
	arr = append(arr, 2, 4, 2, 3, 5, 5, 3, 12, 4, 56, 12, 4, 65, 1, 75, 4, 3, 671234, 634, 234, 6, 45, 2345643, 4525)
	// Ввывод отсортированного массива
	quickSort(arr)
	fmt.Println(arr)

}
