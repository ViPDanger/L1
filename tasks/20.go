//	Разработать программу, которая переворачивает слова в строке.
//	Пример: «snow dog sun — sun dog snow».

package tasks

import (
	"fmt"
	"strings"
)

func wordInvert(str string) string {
	var result strings.Builder
	words := strings.Fields(str)
	// Идём к середине и меняем слова между собой
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}
	// собираем строку result
	for _, word := range words {
		result.WriteString(word)
		result.WriteRune(' ')
	}
	// Возвращаем без лишнего пробела
	return result.String()[:result.Len()-1]
}

func Task20_1() {
	str := "snow dog sun"
	fmt.Println(wordInvert(str))
}
