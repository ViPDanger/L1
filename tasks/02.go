// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
package tasks

import (
	"fmt"
	"sync"
)

func Task2_1() {
	// WaitGroup добавлен для завершения каждого из потоков.
	var wg sync.WaitGroup

	// Массив (2,4,6,8,10)
	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))

	// Возведение в квадраты 7
	for _, a := range arr {
		go func(a int) {
			fmt.Println(a * a)
			wg.Done()
		}(a)
	}
	// Ожидаем выполнения всех горутин
	wg.Wait()
}
