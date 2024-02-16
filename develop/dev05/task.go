package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
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

func main() {
	fmt.Println(
		new(A).highlight("Welcome to Linux !\n"+
			"Linux is a free and opensource Operating system that is mostly used by\n"+
			"developers and in production servers for hosting crucial components such as web\n"+
			"and database servers. Linux has also made a name for itself in PCs.\n"+
			"Beginners looking to experiment with Linux can get started with friendlier linux\n"+
			"distributions such as Ubuntu, Mint, Fedora and Elementary OS.", "Linux", 1),
	)
}

type Highlighter interface {
	highlight(string, string, int) []string
}

// #region -A
type A struct{}

func (a *A) highlight(s string, subStr string, n int) []string {
	lines := strings.Split(s, "\n")
	// subStr = strings.ToLower(subStr)
	var indexies = findAll(lines, subStr)

	var result = make([]string, len(indexies))
	for i, index := range indexies {
		if len(lines)-n >= index {
			result[i] = strings.Join(lines[index:index+n+1], "")
		} else {
			result[i] = strings.Join(lines[index:], "")
		}
	}

	return result
}

// #endregion

func findAll(lines []string, subStr string) (indexies []int) {
	for i, line := range lines {
		if strings.Contains(line, subStr) {
			indexies = append(indexies, i)
		}
	}
	return
}

func findAll2(lines []string, subStr string) (indexies []int) {
	for i, line := range lines {
		if strings.Contains(strings.ToLower(line), subStr) {
			indexies = append(indexies, i)
		}
	}
	return
}

// #region -B
type B struct{}

func (b *B) highlight(s string, subStr string, n int) []string {
	lines := strings.Split(s, "\n")
	// subStr = strings.ToLower(subStr)
	var indexies = findAll(lines, subStr)

	var result = make([]string, len(indexies))
	for i, index := range indexies {
		if index-n > -1 {
			result[i] = strings.Join(lines[index-n:index+1], "")
		} else {
			result[i] = strings.Join(lines[:index+1], "")
		}
	}

	return result
}

// #endregion

// #region -C
type C struct{}

func (c *C) highlight(s string, subStr string, n int) []string {
	lines := strings.Split(s, "\n")
	// subStr = strings.ToLower(subStr)
	var indexies = findAll(lines, subStr)

	var result = make([]string, len(indexies))
	for i, index := range indexies {

		f := 0
		l := len(lines)

		if index-n > -1 {
			f = index - n
		}

		if len(lines)-n >= index {
			f = index + n + 1
		}

		result[i] = strings.Join(lines[f:l], "")
	}

	return result
}

// #endregion

// #region -c
type c struct{}

func (c *c) counter(lines string, subStr string) (count int) {
	for _, line := range strings.Split(lines, "\n") {
		if strings.Contains(line, subStr) {
			count++
		}
	}
	return
}

// #endregion

// #region -i
type I struct{}

func (i *I) highlight(s string, subStr string, n int) []string {
	lines := strings.Split(s, "\n")
	subStr = strings.ToLower(subStr)
	var indexies = findAll(lines, subStr)

	var result = make([]string, len(indexies))
	for i, index := range indexies {
		result[i] = lines[index]
	}

	return result
}

// #endregion

// #region -v
type V struct{}

func (v *V) highlight(s string, subStr string) []string {
	lines := strings.Split(s, "\n")
	subStr = strings.ToLower(subStr)
	var indexies = findAll2(lines, subStr)

	var result = make([]string, 0)
	for i, line := range lines {
		if !slices.Contains(indexies, i) {
			result = append(result, line)
		}
	}

	return result
}

// #endregion

// #region -F
type F struct{}

func (f *F) highlight(s string, subStr string) []int {
	lines := strings.Split(s, "\n")
	result := make([]int, 0)

	for i, line := range lines {
		if line == subStr {
			result = append(result, i)
		}
	}

	return result
}

// #endregion

// #region -n
type N struct{}

func (n *N) find(lines string, subStr string) int {
	for i, v := range strings.Split(lines, "\n") {
		if strings.Contains(v, subStr) {
			return i + 1
		}
	}

	return -1
}

// #endregion
