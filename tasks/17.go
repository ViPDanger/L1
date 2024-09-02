// Реализовать бинарный поиск встроенными методами языка.
package tasks

import "fmt"

func binarySearch(a []int, search int) (result int, searchCount int) {
	wall := len(a) / 2
	switch {
	case len(a) == 0:
		result = -1 // В нашем массиве нет искомого элемента
	case a[wall] > search:
		result, searchCount = binarySearch(a[:wall], search)
	case a[wall] < search:
		result, searchCount = binarySearch(a[wall+1:], search)
		if result >= 0 { // if anything but the -1 "not found" result
			result += wall + 1
		}
	default: // a[mid] == search
		result = wall // found
	}
	searchCount++
	return
}

func Task17_1() {
	// Пример массива с binarysearch
	arr := make([]int, 0)
	arr = append(arr, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6, 12, 12, 45, 56, 65, 75, 234, 634, 4525, 671234, 2345643)
	// Для применения бинарного поиска нужна сортировка по возрастанию!
	quickSort(arr)
	result, _ := binarySearch(arr, 671234)
	fmt.Println("result=", result, ", id=", 671234)
	result, _ = binarySearch(arr, 7)
	fmt.Println("result=", result, ", id=", 7)
}
