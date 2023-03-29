package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func checkAnagram(word1, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	chars := make(map[rune]int)
	for _, v := range word1 {
		chars[v]++
	}
	for _, v := range word2 {
		if _, ok := chars[v]; !ok {
			return false
		}
		if chars[v]-1 == 0 {
			delete(chars, v)
		}
		chars[v]--
	}
	return true
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func GetAnagrams(s []string) map[string][]string {
	res := make(map[string][]string)
	for _, v := range s {
		isIn := false
		for k := range res {

			if checkAnagram(strings.ToLower(k), strings.ToLower(v)) {
				res[k] = append(res[k], strings.ToLower(v))
				isIn = true
			}
		}
		if !isIn {
			res[strings.ToLower(v)] = append(res[strings.ToLower(v)], strings.ToLower(v))

		}
	}
	for k := range res {
		res[k] = unique(res[k])

		if len(res[k]) <= 1 {
			delete(res, k)

		} else {

			sort.Strings(res[k])
		}

	}
	return res
}

func main() {

	strs := []string{"пятак", "пятка", "тяпка", "ПяткА", "листок", "слиток", "столик", "сЛиток", "а", "а", "з"}
	out := GetAnagrams(strs)
	fmt.Println(out)
}
