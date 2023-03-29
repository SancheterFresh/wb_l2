package main

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"errors"
	"fmt"
	"unicode"
)

func UnpackString(s string) (string, error) {

	var result []rune
	var count int

	rns := []rune(s)

	if !unicode.IsDigit(rns[0]) {
		for i := len(rns) - 1; i >= 0; i-- {
			if unicode.IsDigit(rns[i]) {
				count = count*10 + int(rns[i]-'0')
			} else {
				if count == 0 {
					count = 1
				}
				for j := 0; j < count; j++ {
					result = append(result, rns[i])
				}
				count = 0
			}
		}
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}

		return string(result), nil
	} else {
		return "", errors.New("incorrect input")
	}
}

func main() {
	result, err := UnpackString("f4h3g5t3")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
