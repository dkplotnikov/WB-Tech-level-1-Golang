package main

import (
	"fmt"
	"unicode"
)

// Разработать программу, которая проверяет, что все символы
// в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

// IsConsistOfUniqueSymbols возвращает true, если строка состоит
// из уникальных символов, иначе false
func IsConsistOfUniqueSymbols(s string) bool {

	// Множество встретившихся символов
	smblSet := make(map[rune]struct{})
	for _, r := range s {

		// Приведение символа к нижнему регистру
		r = unicode.ToLower(r)

		// Если такой символ уже есть во множестве, возвращаем false
		if _, ok := smblSet[r]; ok {
			return false
		}
		// Если символ ранее не встречался, добавляем его во множество
		smblSet[r] = struct{}{}
	}
	return true
}

func main() {

	str := []string{"abcd", "abCdefAaf", "aabcd", "aA", "  ", ""}
	for _, s := range str {
		fmt.Printf("%s - %v\n", s, IsConsistOfUniqueSymbols(s))
	}

}
