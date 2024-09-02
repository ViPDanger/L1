//	Разработать программу, которая проверяет, что все символы в строке
//	уникальные (true — если уникальные, false etc). Функция проверки должна быть
//	регистронезависимой.

//	Например:
//	abcd — true
//	abCdefAaf — false
//	aabcd — false

package tasks

import (
	"fmt"
	"strings"
)

// Регистронезависимая функция проверки уникальности
func isSymbolsUnique(str string) bool {
	Map := make(map[rune]bool, 0)
	for _, r := range strings.ToLower(str) {
		if Map[r] {
			return false
		}
		Map[r] = true
	}
	return true
}

// Вызов проверочных строк
func Task26_1() {
	fmt.Println("abcd - ", isSymbolsUnique("abcd"))
	fmt.Println("abCdefAaf - ", isSymbolsUnique("abCdefAaf"))
	fmt.Println("aabcd - ", isSymbolsUnique("aabcd"))
}
