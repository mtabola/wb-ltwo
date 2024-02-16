package main

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	testCases := []struct {
		words    []string
		anagrams map[string][]string
	}{
		{
			words: []string{"пятак", "пятка", "тяпка"},
			anagrams: map[string][]string{
				"пятак": {"пятак", "пятка", "тяпка"},
			},
		},
		{
			words: []string{"Просвещение", "Всепрощение", "Воспрещение"},
			anagrams: map[string][]string{
				"просвещение": {"воспрещение", "всепрощение", "просвещение"},
			},
		},
		{
			words: []string{"ПЕСЧАНИК", "ПАСЕЧНИК", "ПЕСЧИНКА"},
			anagrams: map[string][]string{
				"песчаник": {"пасечник", "песчаник", "песчинка"},
			},
		},
		{
			words: []string{"пятак", "пятка", "тяПка", "тяпка", "тЕрка", "терк", "сЛИток", "Листок", "столик", "стил", "лист"},
			anagrams: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"слиток": {"листок", "слиток", "столик"},
				"стил":   {"лист", "стил"},
			},
		},
	}

	for i := range testCases {
		if output := CheckAnagrams(&testCases[i].words); !reflect.DeepEqual(*output, testCases[i].anagrams) {
			t.Errorf("QuickSort does not sort, %v not equal %v", *output, testCases[i].anagrams)
		}
	}

}
