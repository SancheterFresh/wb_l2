package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func commands(commandLine string) {
	command := strings.Split(commandLine, " ")

	switch command[0] {
	case "cd":
		cd := command[1]
		os.Chdir(cd)
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
	case "echo":
		out := strings.Join(command[1:], " ")
		fmt.Println(out)
	case "kill":
		proc := command[1]
		prId, err := strconv.Atoi(proc)
		if err != nil {
			log.Println(err)
		} else {
			process, err := os.FindProcess(prId)
			if err != nil {
				log.Println(err)
			} else {
				process.Kill()
			}
		}
	case "ps":
		porcesses, _ := ps.Processes()
		for _, p := range porcesses {
			fmt.Printf("Process: %v Id: %v\n", p.Executable(), p.Pid())
		}
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("Commend %s is not exists", command[0])
	}
}

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	for scaner.Scan() {
		commands(scaner.Text())
	}
}
