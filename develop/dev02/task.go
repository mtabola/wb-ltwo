package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

func main() {
	line, err := Unpacking("qwe\\\\5")
	fmt.Println(line, err)
}

func Unpacking(line string) (string, error) {

	result := ""
	var slashes = make(map[int]string) // для хранения позиция и типа слешей (двойной/одинарный)

	// цикл для заполнения мапы позиций слешей
	for i2s, i1s := strings.Index(line, "\\\\"), strings.Index(line, "\\"); i2s != -1 || i1s != -1; {
		if i2s != -1 {
			line = strings.Replace(line, "\\\\", "", 1)
			slashes[i2s] = "\\\\"
		} else {
			if i1s != -1 {
				line = strings.Replace(line, "\\", "", 1)
				slashes[i1s] = "\\"
			}
		}

		i2s, i1s = strings.Index(line, "\\\\"), strings.Index(line, "\\")
	}

	// цикл для прохождения по самой строке и преобразования ее в результирующую
	for i := 0; i < len(line); i++ {

		if slashes[i] == "\\\\" {
			result += "\\" // вставка слэша если на этом месте был двойной
		}

		numOfRepeat, err1 := strconv.Atoi(string(line[i])) // парсинг руны в число
		if err1 != nil || slashes[i] == "\\" {
			result += string(line[i])
			continue
		}

		if i == 0 {
			return "", errors.New("некорректная строка") // проверка на то находится ли число в самом начале строки, если да то исходная строка неверная
		}

		_, err2 := strconv.Atoi(string(line[i-1])) // проверка символа идущего до найденого числа
		if err2 == nil && slashes[i-1] == "" {
			return "", errors.New("некорректная строка") // если перед числом стоит другое число => выбрасываем ошибку
		}

		// цикл для вставки n-го количества нужного символа
		for j := 0; j < numOfRepeat-1; j++ {
			if slashes[i] == "\\\\" {
				result += "\\"
			} else {
				result += string(line[i-1])
			}
		}
	}

	return result, nil
}
