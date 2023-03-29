package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flagsList struct {
	f string
	d string
	s bool
}

func setFlags() *flagsList {
	f := new(flagsList)
	flag.StringVar(&f.f, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&f.d, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&f.s, "s", false, "только строки с разделителем")

	flag.Parse()

	return f
}

func main() {

	flags := setFlags()

	scanner := bufio.NewScanner(os.Stdin)

	strs := [][]string{}
	for {
		fmt.Print("Введите строку. Для завершения ввода нажмите enter: ")
		ok := scanner.Scan()
		if !ok && scanner.Err() == nil {
			break
		}
		// Остановка если пустая строка
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		if flags.s {
			if !strings.Contains(str, flags.d) {
				continue
			}

		}
		strs = append(strs, strings.Split(str, flags.d))

	}
	if flags.f != "" {
		fields := strings.Split(flags.f, ",")
		var fids []int
		for _, stfid := range fields {
			fid, err := strconv.Atoi(stfid)
			if err != nil {
				log.Println(err)
			} else {
				if fid < len(strs)+1 {
					fids = append(fids, fid)
				}
			}
		}

		for _, v := range strs {
			var out string
			for _, fid := range fids {

				out += v[fid-1] + "\t"
			}
			fmt.Println(out)
		}

	} else {
		for _, v := range strs {
			fmt.Println(v)
		}
	}

}
