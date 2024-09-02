// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.
package tasks

import (
	"fmt"
	"reflect"
)

// Определяющая функция
func reflectTask14(x interface{}) {
	fmt.Println("Тип переменной - ", reflect.TypeOf(x))
}

func Task14_1() {
	// Определение переменных
	var (
		i int
		s string
		b bool
		c chan int
	)
	// Определение типа данных в рантайме
	reflectTask14(i)
	reflectTask14(s)
	reflectTask14(b)
	reflectTask14(c)
}
