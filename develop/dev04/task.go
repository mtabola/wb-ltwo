/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func deleteDublications(arr []string) []string {
	uniqueCount := make(map[string]bool, 0)
	outArr := make([]string, 0)
	for _, word := range arr {
		if _, ok := uniqueCount[word]; !ok {
			uniqueCount[strings.ToLower(word)] = true
			outArr = append(outArr, strings.ToLower(word))
		}
	}
	return outArr
}

func CheckAnagrams(arr *[]string) *map[string][]string {
	uniqueArr := deleteDublications(*arr)
	uniqueMap := make(map[string][]string)
	outMap := make(map[string][]string)
	for _, word := range uniqueArr {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		key := strings.Join(letters, "")

		if _, ok := uniqueMap[key]; !ok {
			uniqueMap[key] = make([]string, 0)
		}
		uniqueMap[key] = append(uniqueMap[key], word)
	}
	for k := range uniqueMap {
		if len(uniqueMap[k]) > 1 {
			tmpA := append([]string{}, uniqueMap[k]...)
			sort.Strings(tmpA)
			outMap[uniqueMap[k][0]] = tmpA
		}
	}

	return &outMap
}

func main() {
	arr := []string{"пятак", "пятка", "тяПка", "тяпка", "тЕрка", "терк", "Листок", "сЛИток", "столик", "лист", "литс"}
	fmt.Printf("%v\n", CheckAnagrams(&arr))
}
