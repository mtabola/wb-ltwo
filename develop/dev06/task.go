package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
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

func main() {
	fields := make([]int, 0)
	delimiter := "	"
	separated := false

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return
	}

	if i := strings.Index(text, "-f"); i > -1 {
		subStrs := strings.Split(text[i:], " ")

		for _, v := range strings.Split(subStrs[1], ",") {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			fields = append(fields, num)
		}
	}

	if i := strings.Index(text, "-d"); i > -1 {
		subStrs := strings.Split(text[i:], " ")
		if len(subStrs) > 1 {
			delimiter = subStrs[1]
		}
	}

	if strings.Contains(text, "-s") {
		separated = true
	}

	fmt.Println(doSomething(fields, delimiter, separated))
}

func doSomething(fields []int, delimeter string, separated bool) string {
	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	buffer := make([]byte, 1024)

	n, err := file.Read(buffer)

	if n == 0 || err != nil {
		os.Exit(1)
	}

	buffer = bytes.Trim(buffer, "\x00")
	result := ""
	for _, line := range strings.Split(string(buffer), "\n") {
		if separated && !strings.Contains(line, delimeter) {
			result += line
			continue
		}

		for i, v := range strings.Split(line, delimeter) {
			if len(fields) > 0 {
				if !slices.Contains(fields, i+1) {
					continue
				}
			}
			result += v + " "
		}

		result = strings.Trim(result, " ")
		if !strings.HasSuffix(result, "\n") {
			result += "\n"
		}
	}

	return strings.ReplaceAll(result, "\r", "")
}
