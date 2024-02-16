package main

import "fmt"

/*
	Команда — это поведенческий паттерн проектирования, который превращает запросы
	в объекты, позволяя передавать их как аргументы при вызове методов, ставить запросы
	в очередь, логировать их, а также поддерживать отмену операций

	Удобно работать с сигналами передоваемыми в программу

	Преимущества
	 Убирает прямую зависимость между объектами, вызывающими операции и объектами, которые их непосредственно выполняют.
	 Позволяет реализовать простую отмену и повтор операций.
	 Позволяет реализовать отложенный запуск команд.
	 Позволяет собирать сложные команды из простых.
	 Реализует принцип открытости/закрытости.

	Недостатки
	 Усложняет код программы за счёт дополнительных классов.
*/

type CommandI interface {
	Execute() // выполнение
	Undo()    // отмена
}

// Receiver - Получатель
type Radio struct{}

func (r *Radio) On() {
	fmt.Println("Radio on...")
}
func (r *Radio) Off() {
	fmt.Println("Radio off...")
}

// Реализуем интерфейс
type RadioOnCommand struct {
	rr Radio
}

func (r *RadioOnCommand) Execute() {
	r.rr.On()
}
func (r *RadioOnCommand) Undo() {
	r.rr.Off()
}

// Receiver - Получатель
type TV struct{}

func (r *TV) On() {
	fmt.Println("TV on...")
}
func (r *TV) Off() {
	fmt.Println("TV off...")
}

// Реализуем интерфейс
type TVOnCommand struct {
	rr TV
}

func (r *TVOnCommand) Execute() {
	r.rr.On() // передаем сигнал
}
func (r *TVOnCommand) Undo() {
	r.rr.Off()
}

// Invoker - инициатор
type Pult struct {
	comm CommandI
}

func (p *Pult) presButton() {
	p.comm.Execute()
}

func (p *Pult) presUndo() {
	p.comm.Undo()
}

func main() {
	tv := TVOnCommand{}
	pult := Pult{comm: &tv}
	pult.presButton()
	pult.presUndo()
	pult = Pult{comm: &RadioOnCommand{}}
	pult.presButton()
}
