package main

import "fmt"

/*
Фасад - это структурный паттерн, который  предоставляет простой интерфейс к сложной системе
для облегченной работы с подсистемами

Преимущества
	 Изолирует клиентов от компонентов системы.
	 Уменьшает зависимость между подсистемой и клиентами.

Недостатки
	 Фасад рискует стать супреклассом,
		привязанным ко всем классам программы.
*/

type Read struct{}

func (r *Read) readFile(str string) {
	fmt.Printf("Open and read %s\n file", str)
}

type Send struct{}

func (s *Send) pars() {
	fmt.Println("Data parsing...")
}

func (s *Send) sendData() {
	fmt.Println("Sending data...")
}

type Facade struct {
	read *Read
	send *Send
}

// метод вызывает все остальные
// зависимоти для работы приложения
func (f Facade) SendData(str string) {
	facade := f.newFacade()
	facade.read.readFile(str)
	facade.send.pars()
	facade.send.sendData()
}

func (f Facade) newFacade() *Facade {
	return &Facade{
		read: &Read{},
		send: &Send{},
	}
}

func main() {
	f := Facade{}
	f.SendData("Home/dir/my_file")
}
