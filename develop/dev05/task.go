package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flagsList struct {
	after   int
	before  int
	context int
	count   bool
	ignore  bool
	invert  bool
	fixed   bool
	lineNum bool
	query   string
}

func setFlags() *flagsList {
	f := new(flagsList)
	flag.IntVar(&f.after, "A", 0, "print N lines after")
	flag.IntVar(&f.before, "B", 0, "print N lines before")
	flag.IntVar(&f.context, "C", 0, "print N lines before and after")
	flag.BoolVar(&f.count, "c", false, "print lines count")
	flag.BoolVar(&f.ignore, "i", false, "ignore case")
	flag.BoolVar(&f.invert, "v", false, "invert")
	flag.BoolVar(&f.fixed, "F", false, "full string match")
	flag.BoolVar(&f.lineNum, "n", false, "print line number")
	flag.Parse()
	if f.ignore {
		f.query = strings.ToLower(os.Args[len(os.Args)-1])
	} else {
		f.query = os.Args[len(os.Args)-1]
	}
	return f
}

func getInput(flags *flagsList) []string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		if flags.ignore {
			lines = append(lines, strings.ToLower(scanner.Text()))
		} else {
			lines = append(lines, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func printOutput(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getLines(lines []string, flags *flagsList) []string {
	match := []int{}
	if flags.fixed {
		for i, v := range lines {
			if flags.query == v {
				match = append(match, i)
			}
		}
	} else {
		for i, v := range lines {
			isIn, err := regexp.MatchString(flags.query, v)
			if err != nil {
				log.Fatal(err)
			}
			if isIn {
				match = append(match, i)
			}
		}
	}

	var res []string

	if flags.invert {
		for i := range lines {
			if !contains(match, i) {
				if flags.lineNum {
					res = append(res, strconv.Itoa(i+1)+" "+lines[i])
				} else {
					res = append(res, lines[i])
				}
			}
		}
	} else {
		for _, v := range match {
			if flags.context > 0 {
				var start, end int
				if flags.context > v {
					start = 0
				} else {
					start = v - flags.context
				}

				if flags.context > len(lines)-v-1 {
					end = len(lines)
				} else {
					end = v + flags.context + 1
				}
				for i := start; i < end; i++ {
					if flags.lineNum {
						res = append(res, strconv.Itoa(i+1)+" "+lines[i])
					} else {
						res = append(res, lines[i])
					}
				}

			} else {
				if flags.before > 0 {
					if flags.before > v {
						for i := 0; i < v; i++ {
							if flags.lineNum {
								res = append(res, strconv.Itoa(i+1)+" "+lines[i])
							} else {
								res = append(res, lines[i])
							}
						}
					} else {
						for i := v - flags.before; i < v; i++ {
							if flags.lineNum {
								res = append(res, strconv.Itoa(i+1)+" "+lines[i])
							} else {
								res = append(res, lines[i])
							}
						}
					}
				}

				if flags.lineNum {
					res = append(res, ">"+strconv.Itoa(v+1)+" "+lines[v])
				} else {
					res = append(res, ">"+lines[v])
				}

				if flags.after > 0 {
					if flags.after > len(lines)-v-1 {
						for i := v + 1; i < len(lines); i++ {
							if flags.lineNum {
								res = append(res, strconv.Itoa(i+1)+" "+lines[i])
							} else {
								res = append(res, lines[i])
							}
						}
					} else {
						for i := v + 1; i < v+flags.after+1; i++ {
							if flags.lineNum {
								res = append(res, strconv.Itoa(i+1)+" "+lines[i])
							} else {
								res = append(res, lines[i])
							}
						}
					}
				}
			}
			res = append(res, "...")
		}
		if len(res) > 0 {
			res = res[:len(res)-1]
		}
	}
	if flags.count {
		fmt.Printf("Count: %d\n", len(match))
	}

	return res
}

func main() {

	flags := setFlags()
	lines := getInput(flags)
	result := getLines(lines, flags)

	printOutput(result)

}
