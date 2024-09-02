// Поменять местами два числа без создания временной переменной.
package tasks

import (
	"fmt"
	"sync/atomic"
)

//import "sync/atomic"

// Классический вариант
func Task13_1() {
	i1, i2 := 10, 5
	fmt.Println("i1= ", i1, ", i2= ", i2)
	i1, i2 = i2, i1
	fmt.Println("i1= ", i1, ", i2= ", i2)
}

// с использованием пакета atomic
// Быть честным, не уверен, что внутри atomic.SwapInt64 не использует временной переменной
// Но думаю в любом случае вариант интересный. Мы ведь не добавили временной переменной в исполняемой функции
func Task13_2() {
	var i1, i2 int64
	i1, i2 = 10, 5
	fmt.Println("i1= ", i1, ", i2= ", i2)
	i2 = atomic.SwapInt64(&i1, i2)
	fmt.Println("i1= ", i1, ", i2= ", i2)
}
