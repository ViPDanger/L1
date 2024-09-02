// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….)
// с использованием конкурентных вычислений.
package tasks

import (
	"fmt"
	"sync"
)

func Task3_1() {
	var anwser int
	var wg sync.WaitGroup
	// Массив (2,4,6,8,10)
	arr := []int{2, 4, 6, 8, 10}
	wg.Add(len(arr))

	// Запуск конкурентного вычисления для значения "answer"
	for _, a := range arr {
		go func() {
			anwser = anwser + a*a
			wg.Done()
		}()
	}
	// Ожидаем выполнения всех горутин
	wg.Wait()
	fmt.Println(anwser)
}
