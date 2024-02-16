package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

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

func main() {
	fmt.Println(mySort(
		[]string{
			"5M",
			"10K",
			"2M",
			"1K",
		},
		&H{},
	))
}

func mySort(line []string, key Sort, opt ...int) (result []string, err error) {
	return key.sortStrings(line, opt...)
}

type Sort interface {
	sortStrings(line []string, opt ...int) (result []string, err error)
}

// #region -k
type K struct{}

func (k *K) sortStrings(line []string, opt ...int) (result []string, err error) {

	columnNum := opt[0]
	var columns = make(map[string]int)
	var keys = make([]string, 0)

	for i, s := range line {
		key := strings.Split(s, " ")[columnNum-1]
		columns[key] = i
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, k := range keys {
		originalLine := line[columns[k]]
		result = append(result, originalLine)
	}

	return
}

// #endregion

// #region -n
type N struct{}

func (n *N) sortStrings(line []string, opt ...int) (result []string, err error) {

	var numbers = make([]int, len(line))

	for i, s := range line {
		num, err := strconv.Atoi(s)

		if err != nil {
			return []string{}, errors.New("присутствует не целочисленное значчение")
		}
		numbers[i] = num
	}

	sort.Ints(numbers)

	for _, n := range numbers {
		result = append(result, strconv.Itoa(n))
	}

	return
}

// #endregion

// #region -r
type R struct{}

func (r *R) sortStrings(line []string, opt ...int) (result []string, err error) {
	sort.Sort(sort.Reverse(sort.StringSlice(line)))
	return line, nil
}

// #endregion

// #region -u
type U struct{}

func (u *U) sortStrings(line []string, opt ...int) (result []string, err error) {

	var inserted = make(map[string]struct{}, len(line))

	for _, s := range line {
		_, founded := inserted[s]

		if founded {
			continue
		}

		inserted[s] = struct{}{}
		result = append(result, s)
	}

	return
}

// #endregion

// #region -b
type B struct{}

func (b *B) sortStrings(line []string, opt ...int) (result []string, err error) {
	for i := 0; i < len(line); i++ {
		line[i] = strings.TrimSpace(line[i])
	}
	sort.Strings(line)
	return line, nil
}

// #endregion

// #region -c
type C struct{}

func (c *C) sortStrings(line []string, opt ...int) (result []string, err error) {
	if !sort.StringsAreSorted(line) {
		err = errors.New("набор строк не отсортирован")
	}
	return
}

// #endregion

// #region -h
type H struct{}

func (h *H) sortStrings(line []string, opt ...int) (result []string, err error) {

	var reFlightNumbers = regexp.MustCompile("([0-9]+)([KMGTPkmgtp])")
	var m = make(map[string]int, len(line))

	for i, s := range line {
		matches := reFlightNumbers.FindStringSubmatch(s)
		num, err := strconv.Atoi(matches[1])
		if err != nil {
			return []string{}, errors.New("ошибка")
		}

		switch strings.ToLower(matches[2]) {
		case "k":
			num *= 1000
		case "m":
			num *= 1000000
		case "g":
			num *= 1000000000
		case "t":
			num *= 1000000000000
		case "p":
			num *= 1000000000000000
		}

		m[strconv.Itoa(num)] = i
		result = append(result, strconv.Itoa(num))
	}

	sorted, _ := new(N).sortStrings(result)
	for i, s := range sorted {
		result[i] = line[m[s]]
	}

	return

}

// #endregion
