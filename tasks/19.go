// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode

package tasks

import "fmt"

func glavriba(str string) string {
	runeStr := []rune(str)
	// Идём к середине и меняем символы между собой
	for i := 0; i < len(runeStr)/2; i++ {
		runeStr[i], runeStr[len(runeStr)-i-1] = runeStr[len(runeStr)-i-1], runeStr[i]
	}
	return string(runeStr)
}

func Task19_1() {
	str := "глав рыба"
	fmt.Println(glavriba(str))
}
