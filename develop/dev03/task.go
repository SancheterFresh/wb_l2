package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type flagsList struct {
	k int
	n bool
	r bool
	u bool
}

func setFlags() *flagsList {
	f := new(flagsList)
	flag.IntVar(&f.k, "k", -1, "sort by column")
	flag.BoolVar(&f.n, "n", false, "sort by numeric value")
	flag.BoolVar(&f.r, "r", false, "reverse")
	flag.BoolVar(&f.u, "u", false, "unique")
	flag.Parse()
	return f
}

func getInput() []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func writeOutput(s []string, flags flagsList) {
	f, err := os.Create("./output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for i, v := range s {
		if flags.u {
			if i != len(s)-1 {
				if s[i] != s[i+1] {
					w.WriteString(v + "\n")
				}
			} else {
				w.WriteString(v + "\n")
			}
		} else {
			w.WriteString(v + "\n")
		}

	}
	w.Flush()
}

func sortLines(arr []string, flags flagsList) []string {
	sArr := make([]string, len(arr))
	copy(sArr, arr)

	if !flags.n {
		if flags.k >= 0 {
			sort.Slice(sArr, func(i, j int) bool {
				if flags.r {
					return strings.Split(sArr[i], " ")[flags.k] > strings.Split(sArr[j], " ")[flags.k]
				} else {
					return strings.Split(sArr[i], " ")[flags.k] < strings.Split(sArr[j], " ")[flags.k]
				}
			})
		} else {
			sort.Slice(sArr, func(i, j int) bool {
				if flags.r {
					return sArr[i] > sArr[j]

				} else {
					return sArr[i] < sArr[j]
				}
			})
		}
	} else {
		if flags.k < 0 {
			flags.k = 0
		}
		sort.Slice(sArr, func(i, j int) bool {
			anum, _ := strconv.Atoi(strings.Split(sArr[i], " ")[flags.k])
			bnum, _ := strconv.Atoi(strings.Split(sArr[j], " ")[flags.k])

			if flags.r {
				return anum > bnum

			} else {
				return anum < bnum
			}
		})
	}

	return sArr

}

func main() {
	flags := setFlags()

	fmt.Println(flags)

	lines := getInput()

	writeOutput(sortLines(lines, *flags), *flags)

	fmt.Println(sortLines(lines, *flags))
}
