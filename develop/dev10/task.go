package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
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
	var timeout int
	var host string
	flag.IntVar(&timeout, "timeout", 10, "Connection timeout (default 10)")
	flag.StringVar(&host, "host", "127.0.0.1", "Host connection (default 127.0.0.1)")
	flag.Parse()

	_ = telnetServer(time.Duration(timeout) * time.Second)
}

func telnetServer(timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", ":5555", timeout)
	if err != nil {
		fmt.Println(err)
		return err
	}
	go func() {
		for {
			msg, _ := bufio.NewReader(conn).ReadString('\n')
			if msg == "" {
				conn.Close()
				os.Exit(0)
			}
			fmt.Println(msg)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		line, _ := reader.ReadString('\n')
		if strings.Contains(line, string([]byte{4})) {
			conn.Close()
			return nil
		}
		conn.Write([]byte(line))
	}
}
