package main

import (
	"fmt"
	"sync"
	"time"
)

/*
=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func orFn() error {

	var wg sync.WaitGroup
	wg.Add(3)

	or := func(channels ...<-chan time.Time) <-chan time.Time {
		out := make(chan time.Time)

		for _, c := range channels {
			go func(c <-chan time.Time) {
				out <- <-c
				wg.Done()
			}(c)
		}

		go func() {
			wg.Wait()
			close(out)
		}()
		return out
	}

	channels := []<-chan time.Time{time.After(1 * time.Second), time.After(3 * time.Second), time.After(2 * time.Second)}

	for v := range or(channels...) {
		fmt.Println(v)
	}

	return nil
}

func main() {

}
