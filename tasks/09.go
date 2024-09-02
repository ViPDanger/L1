// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout
package tasks

import (
	"fmt"
)

// Буфферизированные каналы можно закрывать до того, как их прочтут (т.к. они буфферизированные). Будем использовать их.

// Первый вариант - с использованием функций для красивого оформления.
func conveyor1(i int) <-chan int {
	result := make(chan int, 1)
	result <- i
	close(result)
	return result
}

func conveyor2(channel <-chan int) <-chan int {
	result := make(chan int, 1)
	numeric := <-channel
	result <- numeric * numeric
	close(result)
	return result
}

func Task9_1() {
	arr := []int{2, 4, 6, 8, 10}
	for _, i := range arr {
		fmt.Println(<-conveyor2(conveyor1(i)))
	}

}

// Во втором варианте создаём каналы прямо внутри основной функции
func pow2(i int) int {
	return i * i
}
func Task9_2() {
	arr := []int{2, 4, 6, 8, 10}
	c := make(chan int, 1)
	cPow2 := make(chan int, 1)
	for _, i := range arr {
		c <- i
		cPow2 <- pow2(<-c)
		fmt.Println(<-cPow2)
	}
	close(c)
	close(cPow2)
}
