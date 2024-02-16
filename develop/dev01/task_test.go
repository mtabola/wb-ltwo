package main

import (
	"fmt"
	"testing"
)

func Test_PrintPresiceTime(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(PrintPresiceTime())
	}
}
