//	К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

//		var justString string
//		func someFunc() {
//		  v := createHugeString(1 << 10)
//		  justString = v[:100]
//		}
//
//		func main() {
//		  someFunc()
//		}

// Комментарий автора
// Во первых, здесь не прооверяется, насколько большую по размеру string выдаёт функция createHugeString
// Мы конечно можем поверить в её названию, но проверку никто не отменял
// Далее, я бы убрал внешнюю переменную justString. Если мы можем просто вернуть слайс функцией, к чему
// нам иметь переменную внутри ПАКЕТА?
// Ну и стоит написать хоть какую то createhugeString. В задании подразумевается, что она где то в пакете есть,
// и передаёт некое значение, однако чисто технически это просто набор букв.
package tasks

import (
	"fmt"
	"strings"
)

// Пример func createHugeString. Сюда можно выдать любую желаемую функцию или интерфейс.
func createHugeString(i int) string {
	var result strings.Builder

	for ; i > 0; i-- {
		result.WriteRune(rune(i))
	}
	return result.String()
}

func someFunc() string {
	v := createHugeString(1 << 10)
	// Проверка на размер полученной строки
	if len([]rune(v)) < 100 {
		return v[:len([]rune(v))]
	}
	return v[:100]
}

func Task15_1() {
	// Просто проверим вывод функции someFunc()
	fmt.Println(someFunc())
}
