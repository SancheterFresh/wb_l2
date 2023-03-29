package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeOut := flag.Int("timeout", 10, "Time out flag")
	flag.Parse()
	args := flag.Args()
	host := args[0]
	port := args[1]

	to := time.Second * time.Duration(*timeOut)

	addr := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", addr, to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connrctng to %s: %s", addr, err)
		os.Exit(1)
	}

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintf(conn, "%s\n", scanner.Text())
		}
		conn.Close()
	}()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	conn.Close()

}
