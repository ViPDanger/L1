// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и
// конструктором.
package tasks

import (
	"fmt"
	"math"
)

//	Комментарий автора
//	Т.К инкапсуляция в Go работает внутри пакета, а по требованию L1
//	каждая задача должна ПОЛНОСТЬЮ содержаться в своём файле,
//	будем моделировать инкапсуляцию не используя инкапсулированные
//	переменные внутри пакета.

type Point struct {
	x float64
	y float64
}

// Сама функция подсчёта расстояния между двумя точками
func (p1 *Point) Range(p2 *Point) float64 {
	return math.Sqrt(math.Pow(p1.x+p2.x, 2) + math.Pow(p1.y+p2.y, 2))
}

// Конструктор Point
func NewPoint(x float64, y float64) *Point {
	return &Point{x, y}
}

func Task24_1() {
	p1 := NewPoint(10., 0.)
	p2 := NewPoint(0., 10.)
	fmt.Println("Range between p1 and p2 =", p1.Range(p2))
}
